package greeting

import (
	"context"
	"encoding/hex"
	"encoding/json"
	internalhttp "github.com/itsnoproblem/frameserver-go/http"
	"github.com/pkg/errors"
	"net/http"
)

type transporter struct {
	makeHandler internalhttp.HandlerMaker
	service     Service
}

func NewTransporter(h internalhttp.HandlerMaker, s Service) *transporter {
	return &transporter{
		makeHandler: h,
		service:     s,
	}
}

func (t *transporter) MakeRoutes() []*internalhttp.Route {
	initialFrameEndpoint := t.makeHandler(
		decodeEmptyRequest,
		makeInitialFrameEndpoint(t.service),
		formatFrameResponse,
	)

	receivePostEndpoint := t.makeHandler(
		decodeReceivePostRequest,
		makeReceivePostEndpoint(t.service),
		formatFrameResponse,
	)

	return []*internalhttp.Route{
		internalhttp.NewRoute(http.MethodGet, "/", initialFrameEndpoint),
		internalhttp.NewRoute(http.MethodPost, "/", receivePostEndpoint),
	}
}

func decodeEmptyRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	return nil, nil
}

type receivePostRequest struct {
	UntrustedData struct {
		FID         int    `json:"fid"`
		URL         string `json:"url"`
		MessageHash string `json:"messageHash"`
		Timestamp   int    `json:"timestamp"`
		Network     int    `json:"network"`
		ButtonIndex int    `json:"buttonIndex"`
		InputText   string `json:"inputText"`
		CastId      struct {
			Fid  int    `json:"fid"`
			Hash string `json:"hash"`
		} `json:"castId"`
	} `json:"untrustedData"`
	TrustedData struct {
		MessageBytes string `json:"messageBytes"`
	} `json:"trustedData"`
}

func decodeReceivePostRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	var req receivePostRequest
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		return nil, errors.Wrap(err, "decodeReceivePostRequest")
	}

	messageBytes, err := hex.DecodeString(req.TrustedData.MessageBytes)
	if err != nil {
		return nil, errors.Wrap(err, "decodeReceivePostRequest")
	}

	return FrameActionRequest{
		MessageBytes: messageBytes,
	}, nil
}
