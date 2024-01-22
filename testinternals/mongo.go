package testinternals

import (
	"context"
	"time"

	"github.com/wearemojo/mojo-public-go/lib/ksuid"
	"go.mongodb.org/mongo-driver/bson"
)

var mongoData = []map[string]any{
	{
		"_id":        ksuid.MustParse("session_000000CmEbzPjqJq0MzT8IflDl9k0"),
		"created_at": time.Now(),

		"device_id": ksuid.MustParse("device_000000CmEbzPjqJq0MzT8IflDl9k1"),
		"platform":  "ios",
		"user_id":   ksuid.MustParse("user_000000CmEbp0hAwaHRb5VSRDWroNl"),
	},
	{
		"_id":        ksuid.MustParse("session_000000CmEbzPjqJq0MzT8IflDl9k2"),
		"created_at": time.Now(),

		"device_id": ksuid.MustParse("device_000000CmEbzPjqJq0MzT8IflDl9k3"),
		"platform":  "ios",
		"user_id":   ksuid.MustParse("user_000000CmEuhLAnUnLPAnzfL0dwK5Q"),
	},
	{
		"_id":        ksuid.MustParse("session_000000CmEbzPjqJq0MzT8IflDl9k4"),
		"created_at": time.Now(),

		"device_id": ksuid.MustParse("device_000000CmEbzPjqJq0MzT8IflDl9k1"),
		"platform":  "ios",
		"user_id":   ksuid.MustParse("user_000000CmEuhLAnUnLPAnzfL0dwK5R"),
	},
	{
		"_id":        ksuid.MustParse("session_000000CmEbzPjqJq0MzT8IflDl9k5"),
		"created_at": time.Now(),

		"device_id": ksuid.MustParse("device_000000CmEbzPjqJq0MzT8IflDl9k1"),
		"platform":  "ios",
		"user_id":   ksuid.MustParse("user_000000CmEbp0hAwaHRb5VSRDWroNo"),
	},
	{
		"_id":        ksuid.MustParse("session_000000CmEbzPjqJq0MzT8IflDl9k6"),
		"created_at": time.Now(),

		"device_id": ksuid.MustParse("device_000000CmEbzPjqJq0MzT8IflDl9k7"),
		"platform":  "ios",
		"user_id":   ksuid.MustParse("user_000000CmEbp0hAwaHRb5VSRDWroNo"),
	},
}

type MongoMock struct{}

func FindAll[T any](_ context.Context, m *MongoMock, filter bson.M) ([]T, error) {
	var filtered []any

outer:
	for idx, record := range mongoData {
		for key, value := range filter {
			if record[key] != value {
				continue outer
			}
		}

		filtered = append(filtered, mongoData[idx])
	}

	out := make([]T, len(filtered))

	for idx, record := range filtered {
		bsonBytes, err := bson.Marshal(record)
		if err != nil {
			return nil, err
		}

		if err := bson.Unmarshal(bsonBytes, &out[idx]); err != nil {
			return nil, err
		}
	}

	return out, nil
}
