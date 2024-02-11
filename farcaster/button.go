package farcaster

const (
	ActionPost         = "post"
	ActionPostRedirect = "post_redirect"
	ActionMint         = "mint"
	ActionLink         = "link"
)

type Button struct {
	// Label (required) a string to be displayed on the button
	Label string
	// Action must be `post`, `post_redirect`, `mint` or `link`. Defaults to `post` if not specified
	Action string
	// Target (optional) URL to send the Signature Packet to.
	Target string
}

// NewPostButton creates a new button with the action set to `post`.
func NewPostButton(label string) Button {
	return Button{
		Label:  label,
		Action: ActionPost,
	}
}

// NewPostRedirectButton creates a new button with the action set to `post_redirect`.
func NewPostRedirectButton(label string, URL string) Button {
	return Button{
		Label:  label,
		Action: ActionPostRedirect,
		Target: URL,
	}
}

// NewMintButton creates a new button with the action set to `mint`.
func NewMintButton(label string) Button {
	return Button{
		Label:  label,
		Action: ActionMint,
	}
}

// NewLinkButton creates a new button with the action set to `link`.
func NewLinkButton(label string, URL string) Button {
	return Button{
		Label:  label,
		Action: ActionLink,
		Target: URL,
	}
}
