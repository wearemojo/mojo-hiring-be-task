package service

import (
	"context"

	"github.com/wearemojo/lead-be-task/session"
	"github.com/xeipuuv/gojsonschema"
)

var ListByDeviceIDSchema = gojsonschema.NewStringLoader(`{
	"type": "object",
	"additionalProperties": false,

	"required": [
		"device_id"
	],

	"properties": {
		"device_id": {
			"type": "string",
			"min_length": 1
		}
	}
}`)

func (s *Service) ListByDeviceID(ctx context.Context, req *session.DeviceID) ([]*session.Session, error) {
	return s.App.ListByDeviceID(ctx, req.DeviceID)
}
