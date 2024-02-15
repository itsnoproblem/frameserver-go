package tile

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"image/color"
)

// Spec is a specification for an image tile.
type Spec struct {
	Text            string
	BackgroundImage string // overrides BackgroundColor if set
	BackgroundColor color.Color
	TextColor       color.Color
	OverlayColor    color.Color
}

// ID returns a unique identifier for the spec.
func (s Spec) ID() string {
	encoded, _ := json.Marshal(s)
	hashed := md5.Sum(encoded)
	return hex.EncodeToString(hashed[:])
}
