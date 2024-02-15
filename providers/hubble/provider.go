package hubble

import (
	"context"

	"github.com/pkg/errors"

	"github.com/itsnoproblem/frameserver-go/farcaster"
)

const (
	pathValidateMessage = "/v1/validateMessage"
)

type HTTPClient interface {
	PostBytes(ctx context.Context, url string, payload []byte, target interface{}) error
}

type provider struct {
	client   HTTPClient
	endpoint string
}

func NewProvider(client HTTPClient, endpoint string) *provider {
	return &provider{
		client:   client,
		endpoint: endpoint,
	}
}

func (p *provider) ValidateMessage(ctx context.Context, message []byte) (farcaster.Message, error) {
	var res ValidateMessageResponse
	url := p.endpoint + pathValidateMessage

	if message == nil {
		return farcaster.Message{}, errors.New("provider.ValidateMessage: message is nil")
	}

	if err := p.client.PostBytes(ctx, url, message, &res); err != nil {
		return farcaster.Message{}, errors.Wrap(err, "provider.ValidateMessage")
	}

	if !res.Valid {
		return farcaster.Message{}, errors.New("provider.ValidateMessage: message is not valid")
	}

	return res.ToFarcasterMessage(), nil
}
