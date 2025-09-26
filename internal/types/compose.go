package internaltypes

import (
	"fmt"
	"time"

	"github.com/compose-spec/compose-go/v2/types"
	"github.com/docker/docker/api/types/container"
)

type Project struct {
	Inspect  *types.Project
	Services *map[string]Service
}

type Service struct {
	Containers *[]container.InspectResponse
}

func (s *Project) NumContainers() int {
	sum := 0
	for _, v := range *s.Services {
		sum += len(*v.Containers)
	}

	return sum
}

func (s *Project) LatestStartedAt() (time.Time, error) {
	var latest time.Time

	for _, c := range *s.Services {
		t, err := c.LatestStartedAt()
		if err != nil {
			return time.Time{}, err
		}

		if t.After(latest) {
			latest = t
		}
	}

	return latest, nil
}

func (s *Service) LatestStartedAt() (time.Time, error) {
	var latest time.Time

	for _, c := range *s.Containers {
		// Parse RFC3339 string from container state
		if c.State == nil || c.State.StartedAt == "" {
			continue
		}

		started, err := time.Parse(time.RFC3339Nano, c.State.StartedAt)
		if err != nil {
			return time.Time{}, err
		}

		if started.After(latest) {
			latest = started
		}
	}

	return latest, nil
}

func (s *Project) Status() string {
	running, exited, restarting, paused := 0, 0, 0, 0

	for _, c := range *s.Services {
		switch c.Status() {
		case "running":
			running++
		case "exited":
			exited++
		case "restarting":
			restarting++
		case "paused":
			paused++
		}
	}

	total := len(*s.Services)

	switch {
	case running == total:
		return "running"
	case exited == total:
		return "exited"
	case restarting > 0:
		return fmt.Sprintf("restarting (%d/%d)", running, total)
	case paused == total:
		return "paused"
	case running > 0:
		return fmt.Sprintf("degraded (%d/%d running)", running, total)
	default:
		return "unknown"
	}
}

func (s *Service) Status() string {
	if len(*s.Containers) == 0 {
		return "No Containers"
	}

	running, exited, restarting, paused := 0, 0, 0, 0

	for _, c := range *s.Containers {
		switch {
		case c.State.Running:
			running++
		case c.State.Dead:
			exited++
		case c.State.Restarting:
			restarting++
		case c.State.Paused:
			paused++
		}
	}

	total := len(*s.Containers)

	switch {
	case running == total:
		return "running"
	case exited == total:
		return "exited"
	case restarting > 0:
		return fmt.Sprintf("restarting (%d/%d)", running, total)
	case paused == total:
		return "paused"
	case running > 0:
		return fmt.Sprintf("degraded (%d/%d running)", running, total)
	default:
		return "unknown"
	}
}
