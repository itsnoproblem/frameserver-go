package greeting

import (
	"context"
	"github.com/itsnoproblem/frameserver-go/templates"
	"github.com/pkg/errors"

	"github.com/a-h/templ"
)

func formatReceivePostResponse(ctx context.Context, response interface{}) (templ.Component, error) {
	return templ.NopComponent, nil
}

func formatFrameResponse(ctx context.Context, response interface{}) (templ.Component, error) {
	res, ok := response.(templates.FrameView)
	if !ok {
		return nil, errors.New("greeting.formatFrameResponse: response is not a FrameView")
	}

	return templates.FarcasterFrame(res), nil
}
