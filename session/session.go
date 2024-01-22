package session

import (
	"context"
	"time"

	"github.com/wearemojo/mojo-public-go/lib/ksuid"
)

type Service interface {
	ListByDeviceID(context.Context, *DeviceID) ([]*Session, error)
	ListByPlatform(context.Context, *Platform) ([]*Session, error)
}

type DeviceID struct {
	DeviceID ksuid.ID `json:"device_id"`
}

type Platform struct {
	Platform string `json:"platform"`
}

type Session struct {
	SessionID ksuid.ID  `json:"session_id"`
	CreatedAt time.Time `json:"created_at"`

	DeviceID ksuid.ID `json:"device_id"`
	Platform string   `json:"platform"`
	UserID   ksuid.ID `json:"user_id"`
}
