package greeting

import (
	"context"

	"github.com/a-h/templ"
	"github.com/pkg/errors"

	"github.com/itsnoproblem/frameserver-go/templates"
)

func formatFrameResponse(ctx context.Context, response interface{}) (templ.Component, error) {
	res, ok := response.(templates.FrameView)
	if !ok {
		return nil, errors.New("greeting.formatFrameResponse: response is not a FrameView")
	}

	return templates.FarcasterFrame(res), nil
}
