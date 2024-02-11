package hubble

import (
	"github.com/itsnoproblem/frameserver-go/farcaster"
	"time"
)

type ValidateMessageResponse struct {
	Valid   bool `json:"valid"`
	Message struct {
		Data struct {
			Type            string `json:"type"`
			FID             int    `json:"fid"`
			Timestamp       int    `json:"timestamp"`
			Network         string `json:"network"`
			FrameActionBody struct {
				Url         string `json:"url"`
				ButtonIndex int    `json:"buttonIndex"`
				InputText   string `json:"inputText"`
				CastID      struct {
					FID  int    `json:"fid"`
					Hash string `json:"hash"`
				} `json:"castId"`
			} `json:"frameActionBody"`
		} `json:"data"`
		Hash            string `json:"hash"`
		HashScheme      string `json:"hashScheme"`
		Signature       string `json:"signature"`
		SignatureScheme string `json:"signatureScheme"`
		Signer          string `json:"signer"`
	} `json:"message"`
}

func (r ValidateMessageResponse) ToFarcasterMessage() farcaster.Message {
	return farcaster.Message{
		Data: farcaster.MessageData{
			Type:      r.Message.Data.Type,
			FID:       r.Message.Data.FID,
			Timestamp: time.Unix(int64(r.Message.Data.Timestamp), 0),
			Network:   r.Message.Data.Network,
			FrameActionBody: farcaster.FrameAction{
				Url:         r.Message.Data.FrameActionBody.Url,
				ButtonIndex: r.Message.Data.FrameActionBody.ButtonIndex,
				InputText:   r.Message.Data.FrameActionBody.InputText,
				CastID: farcaster.CastID{
					FID:  r.Message.Data.FrameActionBody.CastID.FID,
					Hash: r.Message.Data.FrameActionBody.CastID.Hash,
				},
			},
		},
		Hash:            r.Message.Hash,
		HashScheme:      r.Message.HashScheme,
		Signature:       r.Message.Signature,
		SignatureScheme: r.Message.SignatureScheme,
		Signer:          r.Message.Signer,
	}
}
