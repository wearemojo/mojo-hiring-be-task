package app

import (
	"context"

	"github.com/wearemojo/lead-be-task/session"
	"github.com/wearemojo/lead-be-task/session/db"
	"github.com/wearemojo/mojo-public-go/lib/ksuid"
	"github.com/wearemojo/mojo-public-go/lib/slicefn"
)

func (a *App) ListByPlatform(ctx context.Context, platform string) ([]*session.Session, error) {
	dbSessions, err := a.DB.Sessions.ListByPlatform(ctx, platform)
	if err != nil {
		return nil, err
	}

	return mapSessions(dbSessions), nil
}

func (a *App) ListByDeviceID(ctx context.Context, deviceID ksuid.ID) ([]*session.Session, error) {
	dbSessions, err := a.DB.Sessions.ListByDeviceID(ctx, deviceID)
	if err != nil {
		return nil, err
	}

	return mapSessions(dbSessions), nil
}

func mapSessions(dbSessions []*db.Session) []*session.Session {
	return slicefn.Map(dbSessions, func(s *db.Session) *session.Session {
		return &session.Session{
			SessionID: s.ID,
			CreatedAt: s.CreatedAt,

			DeviceID: s.DeviceID,
			Platform: s.Platform,
			UserID:   s.UserID,
		}
	})
}
