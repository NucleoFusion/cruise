package docker

import (
	"context"
	"fmt"

	"github.com/NucleoFusion/cruise/internal/utils"
	"github.com/docker/docker/api/types/volume"
)

func GetNumVolumes() int {
	vols, _ := GetVolumes()
	return len(vols.Volumes)
}

func GetVolumes() (volume.ListResponse, error) {
	return cli.VolumeList(context.Background(), volume.ListOptions{})
}

func VolumesFormattedSummary(vol volume.Volume, width int) string {
	w := width / 11
	size := "N/A"
	if vol.UsageData != nil {
		size = fmt.Sprintf("%dKB", vol.UsageData.Size/1024)
	}
	return fmt.Sprintf("%-*s %-*s %-*s %-*s %-*s",
		3*w, utils.Shorten(vol.Name, 2*w),
		2*w, utils.Shorten(vol.Scope, 2*w),
		2*w, utils.Shorten(vol.Driver, 2*w),
		2*w, utils.Shorten(vol.Mountpoint, 2*w),
		2*w, utils.Shorten(size, 2*w))
}

func VolumesHeaders(width int) string {
	w := width / 11
	return fmt.Sprintf("%-*s %-*s %-*s %-*s %-*s",
		3*w, "Name",
		2*w, "Scope",
		2*w, "Driver",
		2*w, "Mount Point",
		2*w, "Size")
}
