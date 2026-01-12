package dockerruntime

import "github.com/docker/docker/api/types/container"

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
