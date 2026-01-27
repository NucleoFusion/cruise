package docker

import (
	"context"
	"fmt"

	"github.com/cruise-org/cruise/internal/utils"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/volume"
)

func GetNumVolumes() int {
	vols, _ := GetVolumes()
	return len(vols.Volumes)
}

func GetVolumes() (volume.ListResponse, error) {
	return cli.VolumeList(context.Background(), volume.ListOptions{})
}

func InspectVolume(id string) (volume.Volume, error) {
	return cli.VolumeInspect(context.Background(), id)
}

func PruneVolumes() error {
	_, err := cli.VolumesPrune(context.Background(), filters.NewArgs())
	return err
}

func RemoveVolumes(id string) error {
	return cli.VolumeRemove(context.Background(), id, false)
}

func VolumesFormattedSummary(vol volume.Volume, width int) string {
	w := width / 11
	return fmt.Sprintf("%-*s %-*s %-*s %-*s %-*s",
		3*w, utils.Shorten(vol.Name, 2*w),
		1*w, utils.Shorten(vol.Scope, 2*w),
		1*w, utils.Shorten(vol.Driver, 2*w),
		3*w, utils.Shorten(vol.Mountpoint, 2*w),
		3*w, utils.Shorten(vol.CreatedAt, 2*w))
}

func VolumesHeaders(width int) string {
	w := width / 11
	return fmt.Sprintf("%-*s %-*s %-*s %-*s %-*s",
		3*w, "Name",
		1*w, "Scope",
		1*w, "Driver",
		3*w, "Mount Point",
		3*w, "Created")
}
