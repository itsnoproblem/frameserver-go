package farcaster

const (
	AspectRatio_1_1 AspectRatio = "1:1"
	AspectRatio_2_1 AspectRatio = "1.9:1"
)

// Frame represents a frame to be displayed in the farcaster client.
type Frame struct {
	// Version (required) a valid frame version string. The string must be a release date (e.g. 2020-01-01) or vNext.
	Version string
	// Image (required) to be rendered in the frame
	Image Image
	// InputLabel (optional) when set, frame will include an input field with the given label
	InputLabel string
	// Buttons 0 or more buttons to be rendered in the frame
	Buttons []Button
	// PostURL is a 256-byte string which contains a valid URL to send the Signature Packet to
	PostURL string
}
