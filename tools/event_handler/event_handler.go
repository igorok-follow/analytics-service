package event_handler

import (
	"errors"
	"fmt"
	"github.com/igorok-follow/analytics-service/app/models"
	"github.com/igorok-follow/analytics-service/tools/request"
)

type EventHandler struct {
	Queue  chan *models.Event
	Stash  chan *models.Event
	ApiKey string
}

func NewEventHandler(queueLen int, apiKey string) *EventHandler {
	return &EventHandler{
		Queue:  make(chan *models.Event, queueLen),
		Stash:  make(chan *models.Event, queueLen),
		ApiKey: apiKey,
	}
}

func (e *EventHandler) Run() {
	go e.Send()
	go e.Resend()
}

func (e *EventHandler) Add(event *models.Event) error {
	select {
	case e.Queue <- event:
		break
	default:
		return errors.New("queue is crowded")
	}

	return nil
}

func (e *EventHandler) Send() {
	for {
		select {
		case v := <-e.Queue:
			go func(value *models.Event) {
				post, err := request.SendPost("https://api2.amplitude.com/2/httpapi", &models.RegisterEvent{
					ApiKey: e.ApiKey,
					Events: []*models.Event{value},
				})
				if err != nil {
					fmt.Println("event not sent", value, "\nadditional data:", post, "\n\n")
					e.Stash <- value
				}
				fmt.Println("event sent", value, "\nadditional data:", string(post), "\n\n")
			}(v)
		}
	}
}

func (e *EventHandler) Resend() {
	go func() {
		for {
			select {
			case v := <-e.Stash:
				e.Queue <- v
			}
		}
	}()
}
