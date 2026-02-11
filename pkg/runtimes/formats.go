package runtimes

import (
	"fmt"
	"sort"
	"strings"

	"github.com/cruise-org/cruise/internal/utils"
	"github.com/cruise-org/cruise/pkg/types"
)

func ImageFormatted(image types.Image, width int) string {
	name := "<none>:<none>"
	if len(image.Tags) > 0 {
		name = image.Tags[0]
	}

	id := utils.Shorten(strings.TrimPrefix(image.ID, "sha256:"), 20)
	size := utils.FormatSize(image.Size)
	created := utils.CreatedAgo(image.CreatedAt)
	containers := fmt.Sprintf("%d", image.NumContainers)

	format := strings.Repeat(fmt.Sprintf("%%-%ds ", width), 5)

	return fmt.Sprintf(
		format,
		id,
		utils.Shorten(name, width),
		size,
		created,
		containers,
	)
}

func ImageHeaders(width int) string {
	format := strings.Repeat(fmt.Sprintf("%%-%ds ", width), 5)

	return fmt.Sprintf(
		format,
		"ID",
		"RepoTags",
		"Size",
		"Created",
		"Containers",
	)
}

func ContainerFormatted(item types.Container, width int) string {
	format := strings.Repeat(fmt.Sprintf("%%-%ds ", width), 9)

	return fmt.Sprintf(
		format,
		utils.ShortID(item.ID),
		utils.Shorten(item.Name, width),
		utils.Shorten(item.Image, width),
		utils.CreatedAgo(item.Created),
		utils.Shorten(FormatPorts(item.Ports), width),
		item.State,
		utils.Shorten(item.State, width),
		utils.Shorten(FormatMounts(item.Mounts), width),
		utils.Shorten(FormatLabels(item.Labels), width),
	)
}

func ContainerHeaders(width int) string {
	format := strings.Repeat(fmt.Sprintf("%%-%ds ", width), 9)

	return fmt.Sprintf(
		format,
		"ID",
		"Name",
		"Image",
		"Created",
		"Ports",
		"State",
		"Status",
		"Mounts",
		"Size",
	)
}

func FormatPorts(ports []types.ContainerPort) string {
	if len(ports) == 0 {
		return "-"
	}

	var result []string
	for _, p := range ports {
		result = append(result, fmt.Sprintf("%d->%d/%s", p.ContainerPort, p.HostPort, p.Protocol))
	}

	return strings.Join(result, ",")
}

func FormatMounts(mounts []types.ContainerMount) string {
	if len(mounts) == 0 {
		return "-"
	}
	var result []string
	for _, m := range mounts {
		result = append(result, m.Destination)
	}
	return strings.Join(result, ",")
}

func FormatLabels(m map[string]string) string {
	if len(m) == 0 {
		return "-"
	}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	// To achieve consistency
	sort.Strings(keys)

	result := make([]string, 0, len(keys))
	for _, k := range keys {
		result = append(result, fmt.Sprintf("%s=%s", k, m[k]))
	}

	return strings.Join(result, ",")
}
