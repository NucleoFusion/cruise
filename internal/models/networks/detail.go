package networks

import (
	"fmt"
	"time"

	"github.com/NucleoFusion/cruise/internal/styles"
	"github.com/NucleoFusion/cruise/internal/utils"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/docker/docker/api/types/network"
)

type NetworkDetail struct {
	Width   int
	Height  int
	Network network.Summary
}

func NewDetail(w int, h int, ntw network.Summary) *NetworkDetail {
	return &NetworkDetail{
		Width:   w,
		Height:  h,
		Network: ntw,
	}
}

func (s *NetworkDetail) Init() tea.Cmd {
	return nil
}

func (s *NetworkDetail) Update() (*NetworkDetail, tea.Cmd) {
	return s, nil
}

func (s *NetworkDetail) View() string {
	return lipgloss.JoinHorizontal(lipgloss.Center, lipgloss.JoinVertical(lipgloss.Center, s.getDashboardView(), s.getFlagView()),
		lipgloss.JoinVertical(lipgloss.Center, s.getContainerView(),
			lipgloss.JoinHorizontal(lipgloss.Center, s.getOptionsView(), s.getIPAMView())))
}

func (s *NetworkDetail) getDashboardView() string {
	intrn := "✘"
	if s.Network.Internal {
		intrn = "✔"
	}
	ingr := "✘"
	if s.Network.Ingress {
		intrn = "✔"
	}
	text := fmt.Sprintf("%s   %s \n\n%s        %s \n\n%s   %s \n\n%s    %s \n\n%s     %s \n\n%s  %s \n\n%s   %s",
		styles.DetailKeyStyle().Render(" Network: "), styles.TextStyle().Render(s.Network.Name),
		styles.DetailKeyStyle().Render(" ID: "), styles.TextStyle().Render(utils.Shorten(s.Network.ID, s.Width/5)),
		styles.DetailKeyStyle().Render(" Created: "), styles.TextStyle().Render(utils.Shorten(s.Network.Created.Format(time.DateOnly)+" "+s.Network.Created.Format(time.Kitchen), s.Width/5)),
		styles.DetailKeyStyle().Render(" Driver: "), styles.TextStyle().Render(s.Network.Driver),
		styles.DetailKeyStyle().Render(" Scope: "), styles.TextStyle().Render(s.Network.Scope),
		styles.DetailKeyStyle().Render(" Internal: "), styles.TextStyle().Render(intrn),
		styles.DetailKeyStyle().Render(" Ingress: "), styles.TextStyle().Render(ingr))

	return styles.PageStyle().Padding(1, 2).Render(lipgloss.JoinVertical(lipgloss.Center, styles.TitleStyle().Render("Network Details"), "\n\n",
		lipgloss.Place(s.Width/3-6, s.Height*2/3-8, lipgloss.Left, lipgloss.Center, text)))
}

func (s *NetworkDetail) getFlagView() string {
	return styles.PageStyle().Render(lipgloss.Place(s.Width/3-2, s.Height/3-2, lipgloss.Center, lipgloss.Center, "Flags"))
}

func (s *NetworkDetail) getContainerView() string {
	return styles.PageStyle().Render(lipgloss.Place(s.Width*2/3-2, s.Height/3-2, lipgloss.Center, lipgloss.Center, "Container"))
}

func (s *NetworkDetail) getIPAMView() string {
	return styles.PageStyle().Render(lipgloss.Place(s.Width/3-2, s.Height*2/3-2, lipgloss.Center, lipgloss.Center, "IPAM"))
}

func (s *NetworkDetail) getOptionsView() string {
	return styles.PageStyle().Render(lipgloss.Place(s.Width/3-2, s.Height*2/3-2, lipgloss.Center, lipgloss.Center, "Options"))
}
