package db

import (
	"context"

	"github.com/wearemojo/lead-be-task/testinternals"
)

type DB struct {
	Sessions *SessionsCollection
}

func New(ctx context.Context) *DB {
	db := &testinternals.MongoMock{} // DB connection bits would happen here

	return &DB{
		Sessions: newSessionsCollection(ctx, db),
	}
}
