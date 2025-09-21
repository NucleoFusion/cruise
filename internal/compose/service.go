package compose

import (
	"fmt"

	"github.com/NucleoFusion/cruise/internal/types"
	"github.com/NucleoFusion/cruise/internal/utils"
)

func ServiceStatus(s *types.ServiceSummary) string {
	total := 0
	running := 0

	for _, v := range *s.Containers {
		if v.Inspect.State.Running {
			running++
		}

		total++
	}

	var status string
	if total == running {
		status = "Running"
	} else if running != 0 {
		status = "Partially Running"
	} else if running == 0 {
		status = "Exited"
	} else {
		status = "Uknown"
	}

	return status
}

func ServiceHeaders(width int) string {
	w := width / 6
	return fmt.Sprintf("%-*s %-*s %-*s %-*s %-*s %-*s",
		1*w, "Name",
		1*w, "Image",
		1*w, "Status",
		1*w, "CPU",
		1*w, "Mem",
		1*w, "Net",
	)
}

func ServiceFormatted(width int, s *types.ServiceSummary) string {
	w := width / 6
	img := "NA"
	if len(*s.Containers) > 0 {
		img = (*s.Containers)[0].Inspect.Image
	}

	return fmt.Sprintf("%-*s %-*s %-*s %-*d %-*d %-*s",
		1*w, s.Name,
		1*w, utils.ShortID(img),
		1*w, ServiceStatus(s),
		1*w, s.AggregatedStats.CPU,
		1*w, s.AggregatedStats.Mem,
		1*w, fmt.Sprintf("%dRx / %dTx", s.AggregatedStats.NetRx, s.AggregatedStats.NetTx),
	)
}
