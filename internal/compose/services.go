package compose

import "github.com/docker/docker/api/types/container"

type ServiceSummary struct {
	Name       string
	Containers *[]container.Summary
}
