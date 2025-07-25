package data

import "github.com/shirou/gopsutil/v3/mem"

type MemInfo struct {
	Total float64
	Used  float64
	Usage float64
	Err   error
}

func GetMemInfo() *MemInfo {
	v, err := mem.VirtualMemory()

	totalGB := float64(v.Total) / (1 << 30)
	usedGB := float64(v.Used) / (1 << 30)
	usedPercent := v.UsedPercent

	return &MemInfo{
		Total: totalGB,
		Used:  usedGB,
		Usage: usedPercent,
		Err:   err,
	}
}
