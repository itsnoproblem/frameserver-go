package greeting

import (
	"context"

	"github.com/pkg/errors"

	"github.com/itsnoproblem/frameserver-go/farcaster"
	"github.com/itsnoproblem/frameserver-go/tile"
)

type MessageValidator interface {
	ValidateMessage(context.Context, []byte) (farcaster.Message, error)
}

type TileMaker interface {
	MakeTile(spec tile.Spec) (URL string, err error)
}

type service struct {
	validator MessageValidator
	tileMaker TileMaker
	staticDir string
	appURL    string
}

func NewService(validator MessageValidator, tileMaker TileMaker, staticDir, appURL string) *service {
	return &service{
		validator: validator,
		tileMaker: tileMaker,
		staticDir: staticDir,
		appURL:    appURL,
	}
}

func (s *service) ValidateMessage(ctx context.Context, message []byte) (farcaster.Message, error) {
	validatedMessage, err := s.validator.ValidateMessage(ctx, message)
	if err != nil {
		return farcaster.Message{}, errors.Wrap(err, "greeting.ValidateMessage")
	}

	return validatedMessage, nil
}

func (s *service) CreateTile(ctx context.Context, spec tile.Spec) (URL string, err error) {
	tileURL, err := s.tileMaker.MakeTile(spec)
	if err != nil {
		return "", errors.Wrap(err, "greeting.CreateTile")
	}

	return tileURL, nil
}

func (s *service) StaticDir() string {
	return s.staticDir
}

func (s *service) AppURL() string {
	return s.appURL
}
