package networks

import (
	"fmt"
	"strings"
	"time"

	"github.com/NucleoFusion/cruise/internal/styles"
	"github.com/NucleoFusion/cruise/internal/utils"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/docker/docker/api/types/network"
)

type NetworkDetail struct {
	Width    int
	Height   int
	Network  network.Summary
	LabelsVp viewport.Model
	DashVp   viewport.Model
	ContVp   viewport.Model
	IPAMVp   viewport.Model
	OptsVp   viewport.Model
}

func NewDetail(w int, h int, ntw network.Summary) *NetworkDetail {
	labels := make([]string, 0, len(ntw.Labels))
	for k := range ntw.Labels {
		labels = append(labels, k)
	}

	containers := make([]string, 0, len(ntw.Containers))
	for k := range ntw.Containers {
		containers = append(containers, k)
	}

	ipamOpts := make([]string, 0, len(ntw.IPAM.Options))
	for k := range ntw.IPAM.Options {
		ipamOpts = append(ipamOpts, k)
	}

	opts := make([]string, 0, len(ntw.Options))
	for k := range ntw.Options {
		opts = append(opts, k)
	}

	// Label VP
	lvp := viewport.New(w/3, h*2/5)
	lvp.Style = styles.PageStyle().Padding(1, 2)
	lvp.SetContent(getLabelView(ntw, labels, w))

	// Dash VP
	dvp := viewport.New(w/3, h*3/5)
	dvp.Style = styles.PageStyle().Padding(1, 2)
	dvp.SetContent(getDashboardView(ntw, w))

	// Cont VP
	cvp := viewport.New(w*2/3, h/2)
	cvp.Style = styles.PageStyle().Padding(1, 2)
	cvp.SetContent(getContainerView(ntw, containers, w))

	// IPAM VP
	ivp := viewport.New(w/3, h/2)
	ivp.Style = styles.PageStyle().Padding(1, 2)
	ivp.SetContent(getIPAMView(ntw, ipamOpts, w))

	// Options VP
	ovp := viewport.New(w/3, h/2)
	ovp.Style = styles.PageStyle().Padding(1, 2)
	ovp.SetContent(getOptionsView(ntw, opts, w))

	return &NetworkDetail{
		Width:    w,
		Height:   h,
		Network:  ntw,
		LabelsVp: lvp,
		DashVp:   dvp,
		ContVp:   cvp,
		IPAMVp:   ivp,
		OptsVp:   ovp,
	}
}

func (s *NetworkDetail) Init() tea.Cmd {
	return nil
}

func (s *NetworkDetail) Update(msg tea.Msg) (*NetworkDetail, tea.Cmd) {
	return s, nil
}

func (s *NetworkDetail) View() string {
	return lipgloss.JoinHorizontal(lipgloss.Center, lipgloss.JoinVertical(lipgloss.Center, s.DashVp.View(), s.LabelsVp.View()),
		lipgloss.JoinVertical(lipgloss.Center, s.ContVp.View(),
			lipgloss.JoinHorizontal(lipgloss.Center, s.OptsVp.View(), s.IPAMVp.View())))
}

func getDashboardView(ntw network.Summary, w int) string {
	intrn := "✘"
	if ntw.Internal {
		intrn = "✔"
	}
	ingr := "✘"
	if ntw.Ingress {
		intrn = "✔"
	}
	text := fmt.Sprintf("%s   %s \n\n%s        %s \n\n%s   %s \n\n%s    %s \n\n%s     %s \n\n%s  %s \n\n%s   %s",
		styles.DetailKeyStyle().Render(" Network: "), styles.TextStyle().Render(ntw.Name),
		styles.DetailKeyStyle().Render(" ID: "), styles.TextStyle().Render(utils.Shorten(ntw.ID, w/3-15)),
		styles.DetailKeyStyle().Render(" Created: "), styles.TextStyle().Render(utils.Shorten(ntw.Created.Format(time.DateOnly)+" "+ntw.Created.Format(time.Kitchen), w/3-15)),
		styles.DetailKeyStyle().Render(" Driver: "), styles.TextStyle().Render(ntw.Driver),
		styles.DetailKeyStyle().Render(" Scope: "), styles.TextStyle().Render(ntw.Scope),
		styles.DetailKeyStyle().Render(" Internal: "), styles.TextStyle().Render(intrn),
		styles.DetailKeyStyle().Render(" Ingress: "), styles.TextStyle().Render(ingr))

	return lipgloss.JoinVertical(lipgloss.Left, lipgloss.PlaceHorizontal(w/3-4, lipgloss.Center, styles.TitleStyle().Render(" Network Details ")), "\n\n", text)
}

func getContainerView(ntw network.Summary, cntnrs []string, w int) string {
	if len(cntnrs) == 0 {
		return lipgloss.JoinVertical(lipgloss.Center, lipgloss.PlaceHorizontal(w*2/3-4, lipgloss.Center, styles.TitleStyle().Render(" Network Details ")), "\n\n", "No Connected Containers")
	}

	ln := w*2/3 - 2
	text := lipgloss.NewStyle().Bold(true).Render(fmt.Sprintf("%-*s %-*s %-*s %-*s %-*s\n\n",
		ln/5, "ID",
		ln/5, "Name",
		ln/5, "MAC",
		ln/5, "IPv4",
		ln/5, "IPv6",
	))

	for _, v := range cntnrs {
		text += fmt.Sprintf("%-*s %-*s %-*s %-*s %-*s\n\n",
			ln/5, ntw.Containers[v].EndpointID,
			ln/5, ntw.Containers[v].Name,
			ln/5, ntw.Containers[v].MacAddress,
			ln/5, ntw.Containers[v].IPv4Address,
			ln/5, ntw.Containers[v].IPv6Address,
		)
	}

	return lipgloss.JoinVertical(lipgloss.Left, lipgloss.PlaceHorizontal(w*2/3-4, lipgloss.Center, styles.TitleStyle().Render(" Containers ")), "\n\n", text)
}

func getLabelView(ntw network.Summary, labels []string, w int) string {
	text := ""

	if len(labels) == 0 {
		return lipgloss.JoinVertical(lipgloss.Center, lipgloss.PlaceHorizontal(w/3-4, lipgloss.Center, styles.TitleStyle().Render("Labels")), "\n\n", "No Labels Found")
	}

	for _, v := range labels {
		text += fmt.Sprintf("%s %s\n\n", styles.DetailKeyStyle().Render(fmt.Sprintf(" %s: ", utils.Shorten(strings.TrimPrefix(v, "com.docker."), 25))),
			styles.TextStyle().Render(utils.Shorten(ntw.Labels[v], w/3-8-len(v))))
	}

	return lipgloss.JoinVertical(lipgloss.Left, lipgloss.PlaceHorizontal(w/3-4, lipgloss.Center, styles.TitleStyle().Render(" Labels ")), "\n\n", text)
}

func getIPAMView(ntw network.Summary, opts []string, w int) string {
	text := fmt.Sprintf("%s %s\n\n", styles.DetailKeyStyle().Render(fmt.Sprintf(" %s: ", "Driver")),
		styles.TextStyle().Render(ntw.IPAM.Driver))

	for _, v := range opts {
		text += fmt.Sprintf("%s %s\n\n", styles.DetailKeyStyle().Render(fmt.Sprintf(" %s: ", utils.Shorten(v, 25))),
			styles.TextStyle().Render(utils.Shorten(ntw.IPAM.Options[v], w/3-8-len(v))))
	}

	return lipgloss.JoinVertical(lipgloss.Left, lipgloss.PlaceHorizontal(w/3-4, lipgloss.Center, styles.TitleStyle().Render(" IPAM ")), "\n\n", text)
}

func getOptionsView(ntw network.Summary, opts []string, w int) string {
	text := ""

	if len(opts) == 0 {
		return lipgloss.JoinVertical(lipgloss.Center, lipgloss.PlaceHorizontal(w/3-6, lipgloss.Center, styles.TitleStyle().Render(" Options ")), "\n\n", "No Options Found")
	}

	for _, v := range opts {
		text += fmt.Sprintf("%s %s\n\n", styles.DetailKeyStyle().Render(fmt.Sprintf(" %s: ", utils.Shorten(v, 25))),
			styles.TextStyle().Render(utils.Shorten(ntw.Options[v], w/3-8-len(v))))
	}

	return lipgloss.JoinVertical(lipgloss.Left, lipgloss.PlaceHorizontal(w/3-6, lipgloss.Center, styles.TitleStyle().Render(" Options ")), "\n\n", "Check")
}
