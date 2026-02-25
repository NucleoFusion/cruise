package dockerruntime

import (
	"context"
	"time"

	"github.com/cruise-org/cruise/pkg/types"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/api/types/filters"
)

func (s *DockerRuntime) ContainerLogs(ctx context.Context, id string) (*types.Monitor, error) {
	logCh := make(chan types.Log)

	monitor := &types.Monitor{
		Runtime:  "docker",
		Ctx:      ctx,
		Incoming: logCh,
	}

	fltr := filters.NewArgs()
	fltr.Add("container", id)

	msgs, errs := s.Client.Events(ctx, events.ListOptions{
		Filters: fltr,
	})

	go func() {
		defer close(logCh)

		for {
			select {
			case msg, ok := <-msgs:
				if !ok {
					return
				}

				logCh <- types.Log{
					Timestamp: time.Unix(msg.Time, 0),
					Message:   FormatEvent(msg),
				}

			case err, ok := <-errs:
				if ok && err != nil {
					logCh <- types.Log{
						Timestamp: time.Now(),
						Message:   "event error: " + err.Error(),
					}
				}
				return

			case <-ctx.Done():
				return
			}
		}
	}()

	return monitor, nil
}

func (s *DockerRuntime) RuntimeLogs(ctx context.Context) (*types.Monitor, error) {
	logCh := make(chan types.Log)

	monitor := &types.Monitor{
		Runtime:  "docker",
		Ctx:      ctx,
		Incoming: logCh,
	}

	fltr := filters.NewArgs()

	msgs, errs := s.Client.Events(ctx, events.ListOptions{
		Filters: fltr,
	})

	go func() {
		defer close(logCh)

		for {
			select {
			case msg, ok := <-msgs:
				if !ok {
					return
				}
				logCh <- types.Log{
					Timestamp: time.Unix(msg.Time, 0),
					Message:   FormatEvent(msg),
				}

			case err, ok := <-errs:
				if ok && err != nil {
					logCh <- types.Log{
						Timestamp: time.Now(),
						Message:   "event error: " + err.Error(),
					}
				}
				return

			case <-ctx.Done():
				return
			}
		}
	}()

	return monitor, nil
}
