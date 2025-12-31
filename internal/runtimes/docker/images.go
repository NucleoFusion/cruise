package docker

import (
	"context"
	"fmt"
	"strings"

	"github.com/NucleoFusion/cruise/internal/utils"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
)

func GetNumImages() int {
	images, err := cli.ImageList(context.Background(), image.ListOptions{})
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return -1
	}

	return len(images)
}

func GetImages() ([]image.Summary, error) {
	return cli.ImageList(context.Background(), image.ListOptions{All: true})
}

func RemoveImage(id string) error {
	_, err := cli.ImageRemove(context.Background(), id, image.RemoveOptions{PruneChildren: true, Force: false})
	return err
}

func PruneImages() error {
	_, err := cli.ImagesPrune(context.Background(), filters.NewArgs())
	return err
}

func PushImage(img string) error {
	_, err := cli.ImagePush(context.Background(), img, image.PushOptions{})
	return err
}

func PullImage(img string) error {
	_, err := cli.ImagePull(context.Background(), img, image.PullOptions{})
	return err
}

// TODO: Requires a seperate popup
// func BuildImage(img string) error {
// 	_, err := cli.ImageBuild(context.Background(), img, image.PullOptions{})
// 	return err
// }

func ImageHistory(img string) (string, error) {
	layers, err := cli.ImageHistory(context.Background(), img)
	if err != nil {
		return "", err
	}

	text := ""
	for k, v := range layers {
		id := v.ID
		if len(id) > 19 {
			id = id[:19]
		}
		text += fmt.Sprintf("Layer: %-3d %-20s Size: %-7s %s \n", k, id, fmt.Sprintf("%dMB", v.Size/(1024*1024)), v.CreatedBy)
	}

	return text, nil
}

func ImagesHeaders(width int) string {
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

func ImagesFormattedSummary(image image.Summary, width int) string {
	name := "<none>:<none>"
	if len(image.RepoTags) > 0 {
		name = image.RepoTags[0]
	}

	id := utils.Shorten(strings.TrimPrefix(image.ID, "sha256:"), 20)
	size := utils.FormatSize(image.Size)
	created := utils.CreatedAgo(image.Created)
	containers := fmt.Sprintf("%d", image.Containers)

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
