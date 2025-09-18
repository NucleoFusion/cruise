package projects

import (
	"time"

	"github.com/NucleoFusion/cruise/internal/compose"
	"github.com/NucleoFusion/cruise/internal/messages"
	"github.com/NucleoFusion/cruise/internal/types"
	"github.com/NucleoFusion/cruise/internal/utils"
	tea "github.com/charmbracelet/bubbletea"
)

type ProjectDetails struct {
	Width   int
	Height  int
	Summary *types.ProjectSummary
	Inspect *types.Project
}

func NewProjectDetails(w, h int, s *types.ProjectSummary) *ProjectDetails {
	return &ProjectDetails{
		Width:   w,
		Height:  h,
		Summary: s,
	}
}

func (s *ProjectDetails) Init() tea.Cmd {
	return tea.Tick(0, func(_ time.Time) tea.Msg {
		project, err := compose.Inspect(s.Summary)
		if err != nil {
			return utils.ReturnError("Project Details ", "Error Inspecting Project "+s.Summary.Name, err)
		}

		return messages.ProjectInspectResult{
			Project: project,
		}
	})
}

func (s *ProjectDetails) Update(msg tea.Msg) (*ProjectDetails, tea.Cmd) {
	return s, nil
}

func (s *ProjectDetails) View() string {
	return "Details?"
}
