package docker

import (
	"context"
	"fmt"

	"github.com/cruise-org/cruise/internal/utils"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
)

func GetNumNetworks() (int, error) {
	nt, err := GetNetworks()
	return len(nt), err
}

func GetNetworks() ([]network.Summary, error) {
	networks, err := cli.NetworkList(context.Background(), network.ListOptions{})
	if err != nil {
		return networks, err
	}

	return networks, nil
}

func PruneNetworks() error {
	_, err := cli.NetworksPrune(context.Background(), filters.NewArgs())
	return err
}

func RemoveNetwork(id string) error {
	return cli.NetworkRemove(context.Background(), id)
}

func NetworksFormattedSummary(ntwrk network.Summary, width int) string {
	w := width / 14
	ipv := "✘"
	if ntwrk.EnableIPv4 {
		ipv = "✔"
	}
	return fmt.Sprintf("%-*s %-*s %-*s %-*s %-*s %-*s",
		3*w, utils.Shorten(ntwrk.ID, 2*w),
		3*w, utils.Shorten(ntwrk.Name, 3*w),
		2*w, utils.Shorten(ntwrk.Scope, 2*w),
		2*w, utils.Shorten(ntwrk.Driver, 2*w),
		2*w, ipv,
		2*w, fmt.Sprintf("%d", len(ntwrk.Containers)))
}

func NetworksHeaders(width int) string {
	w := width / 14
	return fmt.Sprintf("%-*s %-*s %-*s %-*s %-*s %-*s",
		3*w, "ID",
		3*w, "Name",
		2*w, "Scope",
		2*w, "Driver",
		2*w, "IPv4",
		2*w, "Container Count")
}
