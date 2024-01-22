package db

import (
	"context"
	"time"

	"github.com/wearemojo/lead-be-task/testinternals"
	"github.com/wearemojo/mojo-public-go/lib/ksuid"
	"go.mongodb.org/mongo-driver/bson"
)

type Session struct {
	ID        ksuid.ID  `bson:"_id"`
	CreatedAt time.Time `bson:"created_at"`

	DeviceID ksuid.ID `bson:"device_id"`
	Platform string   `bson:"platform"`
	UserID   ksuid.ID `bson:"user_id"`
}

type SessionsCollection struct {
	col *testinternals.MongoMock
}

func newSessionsCollection(_ context.Context, db *testinternals.MongoMock) *SessionsCollection {
	return &SessionsCollection{db}
}

func (c *SessionsCollection) ListByPlatform(ctx context.Context, platform string) (res []*Session, err error) {
	return testinternals.FindAll[*Session](ctx, c.col, bson.M{"platform": platform})
}

func (c *SessionsCollection) ListByDeviceID(ctx context.Context, deviceID ksuid.ID) (res []*Session, err error) {
	return testinternals.FindAll[*Session](ctx, c.col, bson.M{"device_id": deviceID})
}
