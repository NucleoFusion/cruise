package docker

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/api/types/filters"
)

func RecentEventStream(limit int) chan *events.Message {
	return GetEventStream(limit, filters.NewArgs())
}

func GetEventStream(limit int, fltr filters.Args) chan *events.Message {
	eventChan := make(chan *events.Message, 500)

	go func() {
		msgs, errs := cli.Events(context.Background(), events.ListOptions{Filters: fltr})

		for {
			select {
			case msg := <-msgs:
				eventChan <- &msg
			case err := <-errs:
				if err != nil {
					fmt.Println("Docker Event Error: " + err.Error()) // TODO: Change for Error Popup
				}
			}
		}
	}()

	return eventChan
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
