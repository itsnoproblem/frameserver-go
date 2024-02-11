package farcaster

import "time"

type Message struct {
	Data            MessageData
	Hash            string
	HashScheme      string
	Signature       string
	SignatureScheme string
	Signer          string
}

type MessageData struct {
	Type            string
	FID             int
	Timestamp       time.Time
	Network         string
	FrameActionBody FrameAction
}
