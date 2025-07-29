package containers

import (
	"time"

	"github.com/NucleoFusion/cruise/internal/colors"
	"github.com/NucleoFusion/cruise/internal/docker"
	"github.com/NucleoFusion/cruise/internal/messages"
	"github.com/NucleoFusion/cruise/internal/styles"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/docker/docker/api/types/container"
)

type ContainerList struct {
	Width         int
	Height        int
	Items         []container.Summary
	Err           error
	FilteredItems []container.Summary
	SelectedIndex int
	Ti            textinput.Model
}

func NewContainerList(w int, h int) *ContainerList {
	ti := textinput.New()
	ti.Width = w - 9
	ti.Prompt = " Search: "
	ti.Placeholder = "Press '/' to search..."

	ti.PromptStyle = lipgloss.NewStyle().Foreground(colors.Load().Lavender)
	ti.PlaceholderStyle = lipgloss.NewStyle().Foreground(colors.Load().Surface2)

	return &ContainerList{
		Width:         w,
		Height:        h,
		Ti:            ti,
		SelectedIndex: 0,
	}
}

func (s *ContainerList) Init() tea.Cmd {
	return tea.Tick(0, func(_ time.Time) tea.Msg {
		items, err := docker.GetContainers()
		return messages.ContainerReadyMsg{
			Items: items,
			Err:   err,
		}
	})
}

func (s *ContainerList) Update(msg tea.Msg) (*ContainerList, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.ContainerReadyMsg:
		s.Items = msg.Items
		s.FilteredItems = msg.Items
		s.Err = msg.Err
		return s, nil
	case tea.KeyMsg:
		if s.Ti.Focused() {
			if msg.String() == "esc" {
				s.Ti.Blur()
				return s, nil
			}
			var cmd tea.Cmd
			s.Ti, cmd = s.Ti.Update(msg)
			return s, cmd
		}
		switch msg.String() {
		case "/":
			s.Ti.Focus()
		}
	}
	return s, nil
}

func (s *ContainerList) View() string {
	if s.Err != nil {
		return lipgloss.Place(s.Width, s.Height, lipgloss.Center, lipgloss.Center, "Error: "+s.Err.Error())
	}

	if len(s.Items) == 0 {
		return lipgloss.Place(s.Width, s.Height, lipgloss.Center, lipgloss.Center, "No Containers Found!")
	}

	style := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(colors.Load().Lavender)

	return lipgloss.JoinVertical(lipgloss.Center,
		style.Render(s.Ti.View()),
		style.Padding(1).Foreground(colors.Load().Text).Render(lipgloss.Place(s.Width-1, s.Height-4,
			lipgloss.Center, lipgloss.Top, s.GetListText())))
}

func (s *ContainerList) GetListText() string {
	w := (s.Width)/9 - 1

	text := lipgloss.NewStyle().Bold(true).Render(docker.SummaryHeaders(w)+"\n") + "\n"

	for k, v := range s.FilteredItems {
		line := docker.FormattedSummary(v, w)

		if k == s.SelectedIndex {
			line = lipgloss.NewStyle().Background(colors.Load().Lavender).Foreground(colors.Load().Base).Render(line)
		} else {
			line = styles.TextStyle().Render(line)
		}

		text += line + "\n"
	}

	return text
}
