package docker

import (
	"context"
	"fmt"
	"time"

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
		attrs = attrs[:len(attrs)-1] // trim trailing space
	}

	return fmt.Sprintf("[%s] %s %s %s", eventTime, msg.Type, msg.Action, attrs)
}
