package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/docker/docker/api/types/container"
)

func ShortID(id string) string {
	if len(id) > 12 {
		return id[:12]
	}
	return id
}

func CreatedAgo(ts int64) string {
	return time.Since(time.Unix(ts, 0)).Round(time.Second).String() + " ago"
}

func FormatPorts(ports []container.Port) string {
	if len(ports) == 0 {
		return "-"
	}
	var result []string
	for _, p := range ports {
		if p.PublicPort != 0 {
			result = append(result, fmt.Sprintf("%d->%d/%s", p.PublicPort, p.PrivatePort, p.Type))
		} else {
			result = append(result, fmt.Sprintf("%d/%s", p.PrivatePort, p.Type))
		}
	}
	return strings.Join(result, ",")
}

func FormatMounts(mounts []container.MountPoint) string {
	if len(mounts) == 0 {
		return "-"
	}
	var result []string
	for _, m := range mounts {
		result = append(result, m.Destination)
	}
	return strings.Join(result, ",")
}

func FormatSize(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

func Shorten(s string, max int) string {
	if len(s) <= max {
		return s
	}
	if max <= 3 {
		return s[:max] // no room for "..."
	}
	return s[:max-3] + "..."
}
