package runtimes

import (
	"fmt"
	"sort"
	"strings"

	"github.com/cruise-org/cruise/internal/utils"
	"github.com/cruise-org/cruise/pkg/types"
)

func FormatLog(msg types.Log) string {
	eventTime := msg.Timestamp.Format("15:04:05") // only HH:MM:SS

	return fmt.Sprintf("[%s] %s", eventTime, msg.Message)
}

func VolumeFormatted(vol types.Volume, width int) string {
	w := width / 11
	return fmt.Sprintf("%-*s %-*s %-*s %-*s %-*s",
		3*w, utils.Shorten(vol.Name, 2*w),
		1*w, utils.Shorten(vol.Scope, 2*w),
		1*w, utils.Shorten(vol.Driver, 2*w),
		3*w, utils.Shorten(vol.Mountpoint, 2*w),
		3*w, utils.Shorten(vol.CreatedAt, 2*w))
}

func VolumeHeaders(width int) string {
	w := width / 11
	return fmt.Sprintf("%-*s %-*s %-*s %-*s %-*s",
		3*w, "Name",
		1*w, "Scope",
		1*w, "Driver",
		3*w, "Mount Point",
		3*w, "Created")
}

func NetworkFormatted(ntwrk types.Network, width int) string {
	w := width / 14
	ipv := "✘"
	if ntwrk.IPv4 {
		ipv = "✔"
	}
	return fmt.Sprintf("%-*s %-*s %-*s %-*s %-*s %-*s",
		3*w, utils.Shorten(ntwrk.ID, 2*w),
		3*w, utils.Shorten(ntwrk.Name, 3*w),
		2*w, utils.Shorten(ntwrk.Scope, 2*w),
		2*w, utils.Shorten(ntwrk.Driver, 2*w),
		2*w, ipv,
		2*w, fmt.Sprintf("%d", ntwrk.NumContainers))
}

func NetworkHeaders(width int) string {
	w := width / 14
	return fmt.Sprintf("%-*s %-*s %-*s %-*s %-*s %-*s",
		3*w, "ID",
		3*w, "Name",
		2*w, "Scope",
		2*w, "Driver",
		2*w, "IPv4",
		2*w, "Container Count")
}

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
