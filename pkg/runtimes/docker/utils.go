package dockerruntime

import (
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/events"
)

func GetBlkio(stats container.StatsResponse, platform string) (uint64, uint64) {
	switch platform {
	case "windows":
		return stats.StorageStats.ReadSizeBytes, stats.StorageStats.WriteSizeBytes
	case "linux":
		var read, write uint64
		for _, entry := range stats.BlkioStats.IoServiceBytesRecursive {
			switch entry.Op {
			case "Read":
				read += entry.Value
			case "Write":
				write += entry.Value
			}
		}
		return read, write
	}
	return 0, 0
}

func FormatEvent(e events.Message) string {
	switch e.Action {
	case "start", "die", "stop":
		return fmt.Sprintf(
			"container %s %s",
			e.Actor.Attributes["name"],
			e.Action,
		)
	default:
		return fmt.Sprintf(
			"%s %s",
			e.Type,
			e.Action,
		)
	}
}
