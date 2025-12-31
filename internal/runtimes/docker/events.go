package docker

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/NucleoFusion/cruise/internal/config"
	"github.com/NucleoFusion/cruise/internal/utils"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/api/types/filters"
)

func RecentEventStream(limit int) (<-chan *events.Message, <-chan error) {
	return GetEventStream(limit, filters.NewArgs())
}

func GetEventStream(limit int, fltr filters.Args) (<-chan *events.Message, <-chan error) {
	eventChan := make(chan *events.Message)
	errChan := make(chan error)

	go func() {
		msgs, errs := cli.Events(context.Background(), events.ListOptions{Filters: fltr, Since: "10m"})

		for {
			select {
			case msg := <-msgs:
				eventChan <- &msg
			case err := <-errs:
				errChan <- err
			}
		}
	}()

	return eventChan, errChan
}

func FormatDockerEvent(msg events.Message) string {
	eventTime := time.Unix(msg.Time, 0).Format("15:04:05") // only HH:MM:SS

	// Pick only relevant attributes
	keys := []string{"name", "image", "status"}
	attrs := ""
	for _, key := range keys {
		if val, ok := msg.Actor.Attributes[key]; ok {
			attrs += fmt.Sprintf("%s=%s ", key, val)
		}
	}
	if len(attrs) > 0 {
		attrs = attrs[:len(attrs)-1] // trim
	}

	return fmt.Sprintf("[%s] %s %s %s", eventTime, msg.Type, msg.Action, attrs)
}

func FormatDockerEventVerbose(msg events.Message) string {
	eventTime := time.Unix(msg.Time, 0).Format("15:04:05")

	// useful attr
	var keys []string
	switch msg.Type {
	case "container":
		keys = []string{"name", "image", "exitCode", "signal"}
	case "image":
		keys = []string{"name", "tag"}
	case "network":
		keys = []string{"name", "type"}
	case "volume":
		keys = []string{"name", "driver"}
	case "plugin":
		keys = []string{"name", "type"}
	default:
		keys = []string{} // fallback
	}

	var extras []string
	for _, k := range keys {
		if v, ok := msg.Actor.Attributes[k]; ok && v != "" {
			extras = append(extras, fmt.Sprintf("%s=%s", k, v))
		}
	}

	return fmt.Sprintf(
		"[%s] %-20s %-10s %-10s %s",
		eventTime,
		utils.Shorten(msg.Actor.ID, 20),
		msg.Action,
		msg.Type,
		strings.Join(extras, ", "),
	)
}

func Export(content []string, page string) error {
	filename := fmt.Sprintf("%d:%d_%d-%d_%s", time.Now().Hour(), time.Now().Minute(), time.Now().Day(), time.Now().Month(), page)

	f, err := os.Create(filepath.Join(config.Cfg.Global.ExportDir, filename))
	if err != nil {
		return err
	}

	f.WriteString(strings.Join(content, "\n"))

	return nil
}
