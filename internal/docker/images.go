package docker

import (
	"context"
	"fmt"
	"strings"

	"github.com/NucleoFusion/cruise/internal/utils"
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

	id := utils.ShortID(image.ID)
	size := utils.FormatSize(image.Size)
	created := utils.CreatedAgo(image.Created)
	containers := fmt.Sprintf("%d", image.Containers)

	format := strings.Repeat(fmt.Sprintf("%%-%ds ", width), 5)

	return fmt.Sprintf(
		format,
		utils.Shorten(name, width),
		id,
		created,
		size,
		containers,
	)
}
