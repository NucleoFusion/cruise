package root

import (
	"time"

	"github.com/NucleoFusion/cruise/internal/enums"
	"github.com/NucleoFusion/cruise/internal/messages"
	"github.com/NucleoFusion/cruise/internal/models/containers"
	errorpopup "github.com/NucleoFusion/cruise/internal/models/error"
	"github.com/NucleoFusion/cruise/internal/models/fzf"
	"github.com/NucleoFusion/cruise/internal/models/home"
	"github.com/NucleoFusion/cruise/internal/models/images"
	"github.com/NucleoFusion/cruise/internal/models/monitoring"
	msgpopup "github.com/NucleoFusion/cruise/internal/models/msg"
	"github.com/NucleoFusion/cruise/internal/models/vulnerability"
	tea "github.com/charmbracelet/bubbletea"
	overlay "github.com/rmhubbert/bubbletea-overlay"
)

type Root struct {
	Width       int
	Height      int
	CurrentPage enums.PageType
	PageItems   map[string]enums.PageType
	// Showing Variables
	IsLoading      bool
	IsChangingPage bool
	IsShowingError bool
	IsShowingMsg   bool
	// Pages / Models
	Home          *home.Home
	Containers    *containers.Containers
	Images        *images.Images
	Vulnerability *vulnerability.Vulnerability
	Monitoring    *monitoring.Monitoring
	ErrorPopup    *errorpopup.ErrorPopup
	MsgPopup      *msgpopup.MsgPopup
	PageFzf       fzf.FuzzyFinder
	Overlay       *overlay.Model
}

func NewRoot() *Root {
	return &Root{
		CurrentPage:    enums.Home,
		IsLoading:      true,
		IsShowingError: false,
		PageItems: map[string]enums.PageType{
			"Home":          enums.Home,
			"Containers":    enums.Containers,
			"Images":        enums.Images,
			"Vulnerability": enums.Vulnerability,
			"Monitoring":    enums.Monitoring,
		},
	}
}

func (s *Root) Init() tea.Cmd { return nil }

func (s *Root) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.CloseError:
		s.IsShowingError = false
		return s, nil
	case messages.ErrorMsg:
		s.IsShowingError = true
		s.ErrorPopup = errorpopup.NewErrorPopup(s.Width, s.Height, msg.Msg, msg.Title, msg.Locn)

		var curr tea.Model
		switch s.CurrentPage {
		case enums.Home:
			curr = s.Home
		case enums.Containers:
			curr = s.Containers
		case enums.Images:
			curr = s.Images
		case enums.Vulnerability:
			curr = s.Vulnerability
		case enums.Monitoring:
			curr = s.Monitoring
		}

		s.Overlay = overlay.New(s.ErrorPopup, curr, overlay.Right, overlay.Top, 2, 2)
		return s, tea.Tick(3*time.Second, func(_ time.Time) tea.Msg { return messages.CloseError{} })
	case messages.CloseMsgPopup:
		s.IsShowingMsg = false
		return s, nil
	case messages.MsgPopup:
		s.IsShowingMsg = true
		s.MsgPopup = msgpopup.NewMsgPopup(s.Width, s.Height, msg.Msg, msg.Title, msg.Locn)

		var curr tea.Model
		switch s.CurrentPage {
		case enums.Home:
			curr = s.Home
		case enums.Containers:
			curr = s.Containers
		case enums.Images:
			curr = s.Images
		case enums.Vulnerability:
			curr = s.Images
		case enums.Monitoring:
			curr = s.Monitoring
		}

		s.Overlay = overlay.New(s.MsgPopup, curr, overlay.Right, overlay.Top, 2, 2)
		return s, tea.Tick(3*time.Second, func(_ time.Time) tea.Msg { return messages.CloseMsgPopup{} })
	case messages.ContainerReadyMsg:
		cnt, cmd := s.Containers.Update(msg)
		s.Containers = cnt.(*containers.Containers)
		return s, cmd
	case messages.ImagesReadyMsg:
		img, cmd := s.Images.Update(msg)
		s.Images = img.(*images.Images)
		return s, cmd
	case messages.FzfSelection:
		s.CurrentPage = s.PageItems[msg.Selection]
		s.IsChangingPage = false
		var cmd tea.Cmd
		switch s.CurrentPage {
		case enums.Home:
			cmd = s.Home.Init()
		case enums.Containers:
			cmd = s.Containers.Init()
		case enums.Images:
			cmd = s.Images.Init()
		case enums.Vulnerability:
			cmd = s.Vulnerability.Init()
		case enums.Monitoring:
			cmd = s.Monitoring.Init()
		}
		return s, cmd
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return s, tea.Quit
		case tea.KeyTab:
			s.IsChangingPage = true
			return s, nil
		}
	case tea.WindowSizeMsg:
		s.Width = msg.Width
		s.Height = msg.Height

		s.PageFzf = fzf.NewFzf([]string{"Home", "Containers", "Images", "Vulnerability", "Monitoring"}, msg.Width, msg.Height)
		s.Home = home.NewHome(msg.Width, msg.Height)
		s.Containers = containers.NewContainers(msg.Width, msg.Height)
		s.Images = images.NewImages(msg.Width, msg.Height)
		s.Vulnerability = vulnerability.NewVulnerability(msg.Width, msg.Height)
		s.Monitoring = monitoring.NewMonitoring(msg.Width, msg.Height)

		cmd := tea.Batch(s.Home.Init(), s.Containers.Init(), s.Images.Init())

		s.IsLoading = false
		return s, cmd
	}

	if s.IsChangingPage {
		var cmd tea.Cmd
		s.PageFzf, cmd = s.PageFzf.Update(msg)
		return s, cmd
	}

	switch s.CurrentPage {
	case enums.Home:
		m, cmd := s.Home.Update(msg)
		s.Home = m.(*home.Home)
		return s, cmd
	case enums.Containers:
		cnt, cmd := s.Containers.Update(msg)
		s.Containers = cnt.(*containers.Containers)
		return s, cmd
	case enums.Images:
		img, cmd := s.Images.Update(msg)
		s.Images = img.(*images.Images)
		return s, cmd
	case enums.Vulnerability:
		img, cmd := s.Vulnerability.Update(msg)
		s.Vulnerability = img.(*vulnerability.Vulnerability)
		return s, cmd
	case enums.Monitoring:
		img, cmd := s.Monitoring.Update(msg)
		s.Monitoring = img.(*monitoring.Monitoring)
		return s, cmd
	}

	return s, nil
}

func (s *Root) View() string {
	if s.IsLoading {
		return "\nLoading..."
	}

	if s.IsShowingError || s.IsShowingMsg {
		return s.Overlay.View()
	}

	if s.IsChangingPage {
		return s.PageFzf.View()
	}

	switch s.CurrentPage {
	case enums.Home:
		return s.Home.View()
	case enums.Containers:
		return s.Containers.View()
	case enums.Images:
		return s.Images.View()
	case enums.Vulnerability:
		return s.Vulnerability.View()
	case enums.Monitoring:
		return s.Monitoring.View()
	}

	return "Cruise - A TUI Docker Client"
}
