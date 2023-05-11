package helpers

import (
	"errors"
	"github.com/igorok-follow/analytics-service/extra/api"
)

func ValidateRegisterEvent(in *api.RegisterEventReq) error {
	switch {
	case in.EventType == "":
		return errors.New("event type is empty")
	}

	return nil
}
