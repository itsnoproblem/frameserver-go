package farcaster

import "time"

// Message represents a message from farcaster to the frame server
type Message struct {
	Data            MessageData
	Hash            string
	HashScheme      string
	Signature       string
	SignatureScheme string
	Signer          string
}

// MessageData holds the actual data of a message
type MessageData struct {
	Type            string
	FID             int
	Timestamp       time.Time
	Network         string
	FrameActionBody FrameAction
}

// FrameAction represents an action performed within a frame in the farcaster client.
type FrameAction struct {
	Url         string
	ButtonIndex int
	InputText   string
	CastID      CastID
}

// CastID represents a farcaster account
type CastID struct {
	FID  int
	Hash string
}
