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

func CalculateCPUPercent(stats container.StatsResponse) float64 {
	cpuDelta := float64(stats.CPUStats.CPUUsage.TotalUsage - stats.PreCPUStats.CPUUsage.TotalUsage)
	systemDelta := float64(stats.CPUStats.SystemUsage - stats.PreCPUStats.SystemUsage)
	onlineCPUs := float64(stats.CPUStats.OnlineCPUs)
	if onlineCPUs == 0 {
		onlineCPUs = float64(len(stats.CPUStats.CPUUsage.PercpuUsage)) // Fallback
	}

	if systemDelta > 0.0 && cpuDelta > 0.0 {
		return (cpuDelta / systemDelta) * onlineCPUs * 100.0
	}
	return 0.0
}

func CalculateMemoryPercent(stats container.StatsResponse) float64 {
	used := float64(stats.MemoryStats.Usage - stats.MemoryStats.Stats["cache"])
	total := float64(stats.MemoryStats.Limit)

	percent := (used / total) * 100.0

	return percent
}

// Wraps the text to the given length and also limits no. of lines, adds a "..." line if exceeding.
func WrapAndLimit(s string, maxLen, maxLines int) string {
	var lines []string

	for i := 0; i < len(s); i += maxLen {
		end := i + maxLen
		if end > len(s) {
			end = len(s)
		}
		lines = append(lines, s[i:end])
	}

	if len(lines) > maxLines {
		lines = append(lines[:maxLines], "...")
	}

	return strings.Join(lines, "\n")
}
