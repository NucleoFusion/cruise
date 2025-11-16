package volumes

import (
	"fmt"
	"strings"

	"github.com/NucleoFusion/cruise/internal/docker"
	"github.com/NucleoFusion/cruise/internal/styles"
	"github.com/NucleoFusion/cruise/internal/utils"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/docker/docker/api/types/volume"
)

type VolumeDetail struct {
	Width    int
	Height   int
	Volume   volume.Volume
	LabelsVp viewport.Model
	DashVp   viewport.Model
	UsageVp  viewport.Model
	StatusVp viewport.Model
	OptsVp   viewport.Model
}

func NewDetail(w int, h int, vol volume.Volume) *VolumeDetail {
	v, err := docker.InspectVolume(vol.Name)
	if err != nil {
		v = vol
	}

	labels := make([]string, 0, len(v.Labels))
	for k := range v.Labels {
		labels = append(labels, k)
	}

	opts := make([]string, 0, len(v.Options))
	for k := range v.Options {
		opts = append(opts, k)
	}

	status := make([]string, 0, len(v.Status))
	for k := range v.Status {
		opts = append(opts, k)
	}

	// Label VP
	lvp := viewport.New((w-2)/3, h-h*3/5)
	lvp.Style = styles.PageStyle().Padding(1, 2)
	lvp.SetContent(getLabelView(v, labels, (w-2)/3-4))

	// Dash VP
	dvp := viewport.New((w-2)/3, h*3/5)
	dvp.Style = styles.PageStyle().Padding(1, 2)
	dvp.SetContent(getDashboardView(v, (w-2)/3-4))

	// Status VP
	svp := viewport.New(w-2-(w-2)*2/3, h)
	svp.Style = styles.PageStyle().Padding(1, 2)
	svp.SetContent(getStatusView(v, status, w-2-(w-2)*2/3-4))

	// Usage VP
	uvp := viewport.New((w-2)/3, h/2)
	uvp.Style = styles.PageStyle().Padding(1, 2)
	uvp.SetContent(getUsageView(v, (w-2)/3-4))

	// Options VP
	ovp := viewport.New((w-2)/3, h-h/2)
	ovp.Style = styles.PageStyle().Padding(1, 2)
	ovp.SetContent(getOptionsView(v, opts, (w-2)/3-4))

	return &VolumeDetail{
		Width:    w,
		Height:   h,
		Volume:   v,
		LabelsVp: lvp,
		DashVp:   dvp,
		StatusVp: svp,
		UsageVp:  uvp,
		OptsVp:   ovp,
	}
}

func (s *VolumeDetail) Init() tea.Cmd {
	return nil
}

func (s *VolumeDetail) Update(msg tea.Msg) (*VolumeDetail, tea.Cmd) {
	return s, nil
}

func (s *VolumeDetail) View() string {
	return lipgloss.JoinHorizontal(lipgloss.Center, lipgloss.JoinVertical(lipgloss.Center, s.DashVp.View(), s.LabelsVp.View()),
		lipgloss.JoinVertical(lipgloss.Center, s.UsageVp.View(), s.OptsVp.View()), s.StatusVp.View())
}

func getDashboardView(vol volume.Volume, w int) string {
	text := fmt.Sprintf("%s %s \n\n%s %s \n\n%s %s \n\n%s %s \n\n%s %s",
		styles.DetailKeyStyle().Render(" Name: "), styles.TextStyle().Render(vol.Name),
		styles.DetailKeyStyle().Render(" Scope: "), styles.TextStyle().Render(utils.Shorten(vol.Scope, w-10)),
		styles.DetailKeyStyle().Render(" Driver: "), styles.TextStyle().Render(vol.Driver),
		styles.DetailKeyStyle().Render(" MountPoint: "), styles.TextStyle().Render(utils.Shorten(vol.Mountpoint, w-10)),
		styles.DetailKeyStyle().Render(" Created: "), styles.TextStyle().Render(utils.Shorten(vol.CreatedAt, w-10)))

	return lipgloss.JoinVertical(lipgloss.Left, lipgloss.PlaceHorizontal(w, lipgloss.Center, styles.TitleStyle().Render(" Volume Details ")), "\n\n", text)
}

func getUsageView(vol volume.Volume, w int) string {
	if vol.UsageData == nil {
		return lipgloss.JoinVertical(lipgloss.Center, lipgloss.PlaceHorizontal(w, lipgloss.Center, styles.TitleStyle().Render(" Usage ")), "\n\n", "NA")
	}

	text := fmt.Sprintf("%s %s \n\n%s %sKb ",
		styles.DetailKeyStyle().Render(" RefCount: "), styles.TextStyle().Render(fmt.Sprintf("%d", vol.UsageData.RefCount)),
		styles.DetailKeyStyle().Render(" Size: "), styles.TextStyle().Render(fmt.Sprintf("%d", vol.UsageData.Size/1024)))
	return lipgloss.JoinVertical(lipgloss.Left, lipgloss.PlaceHorizontal(w, lipgloss.Center, styles.TitleStyle().Render(" Usage ")), "\n\n", text)
}

func getLabelView(vol volume.Volume, labels []string, w int) string {
	text := ""

	if len(labels) == 0 {
		return lipgloss.JoinVertical(lipgloss.Center, lipgloss.PlaceHorizontal(w, lipgloss.Center, styles.TitleStyle().Render("Labels")), "\n\n", "No Labels Found")
	}

	for _, v := range labels {
		text += fmt.Sprintf("%s %s\n\n", styles.DetailKeyStyle().Render(fmt.Sprintf(" %s: ", utils.Shorten(strings.TrimPrefix(v, "com.docker."), 25))),
			styles.TextStyle().Render(utils.Shorten(vol.Labels[v], w-8-len(v))))
	}

	return lipgloss.JoinVertical(lipgloss.Left, lipgloss.PlaceHorizontal(w, lipgloss.Center, styles.TitleStyle().Render(" Labels ")), "\n\n", text)
}

func getStatusView(vol volume.Volume, opts []string, w int) string {
	text := ""

	if len(opts) == 0 {
		return lipgloss.JoinVertical(lipgloss.Center, lipgloss.PlaceHorizontal(w, lipgloss.Center, styles.TitleStyle().Render(" Status ")), "\n\n", "No Status Found")
	}

	for _, v := range opts {
		text += fmt.Sprintf("%s %s\n\n", styles.DetailKeyStyle().Render(fmt.Sprintf(" %s: ", v)),
			styles.TextStyle().Render(utils.Shorten(vol.Options[v], w-8-len(v))))
	}

	return lipgloss.JoinVertical(lipgloss.Left, lipgloss.PlaceHorizontal(w, lipgloss.Center, styles.TitleStyle().Render(" Status ")), "\n\n", text)
}

func getOptionsView(vol volume.Volume, opts []string, w int) string {
	text := ""

	if len(opts) == 0 {
		return lipgloss.JoinVertical(lipgloss.Center, lipgloss.PlaceHorizontal(w, lipgloss.Center, styles.TitleStyle().Render(" Options ")), "\n\n", "No Options Found")
	}

	for _, v := range opts {
		text += fmt.Sprintf("%s %s\n\n", styles.DetailKeyStyle().Render(fmt.Sprintf(" %s: ", utils.Shorten(v, 25))),
			styles.TextStyle().Render(utils.Shorten(vol.Options[v], w-8-len(v))))
	}

	return lipgloss.JoinVertical(lipgloss.Left, lipgloss.PlaceHorizontal(w, lipgloss.Center, styles.TitleStyle().Render(" Options ")), "\n\n", "Check")
}
