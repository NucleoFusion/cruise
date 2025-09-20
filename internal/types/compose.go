package types

import (
	"encoding/json"

	"github.com/docker/docker/api/types/container"
)

type Project struct {
	Name            string
	Services        map[string]*ServiceSummary
	AggregatedStats AggregatedStats
}

type ProjectSummary struct {
	Name               string
	Containers         int
	Services           map[string]bool
	Volumes            int
	Networks           int
	RegistryConfigured bool
}

type ServiceSummary struct {
	Name            string
	Containers      *[]ServiceContainer
	AggregatedStats AggregatedStats
}

type AggregatedStats struct {
	CPU      uint64
	Mem      uint64
	MemLimit uint64
	NetRx    uint64
	NetTx    uint64
	BlkRead  uint64
	BlkWrite uint64
}

type ServiceContainer struct {
	Inspect container.InspectResponse
	Stats   *container.StatsResponseReader
	Decoder *json.Decoder
}

func (s *Project) AggregateStats() error {
	aggr := AggregatedStats{}
	t := uint64(len(s.Services))
	for _, v := range s.Services {
		v.AggregateStats()

		aggr.CPU += v.AggregatedStats.CPU / t // Taking divide by t(len) to avoid overflow
		aggr.Mem += v.AggregatedStats.Mem / t
		aggr.MemLimit += v.AggregatedStats.MemLimit / t
		aggr.NetRx += v.AggregatedStats.NetRx / t
		aggr.NetTx += v.AggregatedStats.NetTx / t
		aggr.BlkRead += v.AggregatedStats.BlkRead / t
		aggr.BlkWrite += v.AggregatedStats.BlkWrite / t
	}

	s.AggregatedStats = aggr

	return nil
}

func (s *ServiceSummary) AggregateStats() error {
	aggr := AggregatedStats{}
	t := uint64(len(*s.Containers))
	for _, v := range *s.Containers {
		var stats container.StatsResponse
		err := v.Decoder.Decode(&stats)
		if err != nil {
			return err
		}

		var rx, tx uint64
		for _, net := range stats.Networks {
			rx += net.RxBytes
			rx += net.TxBytes
		}

		var readBytes, writeBytes uint64
		for _, entry := range stats.BlkioStats.IoServiceBytesRecursive {
			switch entry.Op {
			case "Read":
				readBytes += entry.Value
			case "Write":
				writeBytes += entry.Value
			}
		}

		aggr.CPU += stats.CPUStats.CPUUsage.TotalUsage / t // Taking divide by t(len) to avoid overflow
		aggr.Mem += stats.MemoryStats.Usage / t
		aggr.MemLimit += stats.MemoryStats.Limit / t
		aggr.NetRx += rx / t
		aggr.NetTx += tx / t
		aggr.BlkRead += readBytes / t
		aggr.BlkWrite += writeBytes / t
	}

	s.AggregatedStats = aggr

	return nil
}

func (s *ProjectSummary) NumServices() int {
	sum := 0
	for range s.Services {
		sum += 1
	}

	return sum
}
