package service

import (
	"context"

	"github.com/wearemojo/lead-be-task/session"
	"github.com/xeipuuv/gojsonschema"
)

var ListByPlatformSchema = gojsonschema.NewStringLoader(`{
	"type": "object",
	"additionalProperties": false,

	"required": [
		"platform"
	],

	"properties": {
		"platform": {
			"type": "string",
			"enum": ["ios", "android", "web"]
		}
	}
}`)

func (s *Service) ListByPlatform(ctx context.Context, req *session.Platform) ([]*session.Session, error) {
	return s.App.ListByPlatform(ctx, req.Platform)
}
