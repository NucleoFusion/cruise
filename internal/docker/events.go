package docker

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/api/types/filters"
)

type EventsInfo struct {
	Events []events.Message
}

// func GetRecentEvents(limit int) *EventsInfo {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
//
// 	since := time.Now().Add(-2 * time.Minute).Format(time.RFC3339)
//
// 	options := events.ListOptions{
// 		Filters: filters.NewArgs(),
// 		Since:   since,
// 	}
//
// 	msgs, errs := cli.Events(ctx, options)
// 	events := make([]events.Message, 0, limit)
//
// 	// Will also break if timeout reached
// 	for len(events) < limit {
// 		select {
// 		case msg := <-msgs:
// 			events = append(events, msg)
// 		case err := <-errs:
// 			if err != nil && err != context.Canceled {
// 				fmt.Println("Error:", err)
// 			}
// 			return &EventsInfo{Events: events}
// 		case <-ctx.Done():
// 			return &EventsInfo{Events: events}
// 		}
// 	}
//
// 	return &EventsInfo{Events: events}
// }

func RecentEventStream(limit int) chan *events.Message {
	eventChan := make(chan *events.Message, 500)

	go func() {
		msgs, errs := cli.Events(context.Background(), events.ListOptions{Filters: filters.NewArgs()})

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
