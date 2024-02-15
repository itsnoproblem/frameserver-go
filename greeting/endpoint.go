package greeting

import (
	"context"
	"fmt"
	"image/color"

	"github.com/pkg/errors"

	"github.com/itsnoproblem/frameserver-go/farcaster"
	internalhttp "github.com/itsnoproblem/frameserver-go/http"
	"github.com/itsnoproblem/frameserver-go/templates"
	"github.com/itsnoproblem/frameserver-go/tile"
)

type Service interface {
	ValidateMessage(ctx context.Context, messageBytes []byte) (farcaster.Message, error)
	CreateTile(ctx context.Context, spec tile.Spec) (string, error)
	StaticDir() string
	AppURL() string
}

type FrameActionRequest struct {
	MessageBytes []byte
}

func makeInitialFrameEndpoint(svc Service) internalhttp.HandlerFunc {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		postButton := farcaster.NewPostButton("reveal")
		postButton.Target = svc.AppURL() + "/target" // what even is a target?
		frameImage, err := svc.CreateTile(ctx, tile.Spec{
			Text:            "gm farcaster!",
			BackgroundImage: svc.StaticDir() + "/background.jpg",
			TextColor:       color.White,
			OverlayColor:    color.RGBA{R: 0, G: 0, B: 0, A: 80},
		})
		if err != nil {
			return nil, errors.Wrap(err, "makeInitialFrameEndpoint")
		}

		return templates.FrameView{
			Title: "gm farcaster!",
			Frame: farcaster.Frame{
				Version: "vNext",
				PostURL: svc.AppURL() + "/post",
				Image: farcaster.Image{
					URL:         frameImage,
					AspectRatio: farcaster.AspectRatio_2_1,
				},
				Buttons: []farcaster.Button{
					postButton,
				},
			},
		}, nil
	}
}

func makeReceivePostEndpoint(svc Service) internalhttp.HandlerFunc {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(FrameActionRequest)
		if !ok {
			return nil, errors.New("makeReceivePostEndpoint: failed to parse request")
		}

		msg, err := svc.ValidateMessage(ctx, req.MessageBytes)
		if err != nil {
			return nil, errors.Wrap(err, "makeReceivePostEndpoint")
		}

		frameImage, err := svc.CreateTile(ctx, tile.Spec{
			Text:            fmt.Sprintf("got a verified click from FID %d", msg.Data.FID),
			BackgroundImage: "https://tile.loc.gov/storage-services/service/pnp/cph/3g00000/3g06000/3g06400/3g06458r.jpg",
			TextColor:       color.White,
			OverlayColor:    color.RGBA{R: 0, G: 0, B: 0, A: 105},
		})
		if err != nil {
			return nil, errors.Wrap(err, "makeReceivePostEndpoint")
		}

		redirectButton := farcaster.NewLinkButton("view the code", "https://github.com/itsnoproblem/frameserver-go")

		return templates.FrameView{
			Frame: farcaster.Frame{
				Version: "vNext",
				Image: farcaster.Image{
					URL:         frameImage,
					AspectRatio: farcaster.AspectRatio_2_1,
				},
				Buttons: []farcaster.Button{
					redirectButton,
				},
			},
		}, nil
	}
}
