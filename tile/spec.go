package tile

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"image/color"
)

type Spec struct {
	Text            string
	BackgroundImage string
	TextColor       color.Color
	BackgroundColor color.Color
	OverlayColor    color.Color
}

func (s Spec) ID() string {
	encoded, _ := json.Marshal(s)
	hashed := md5.Sum(encoded)
	return hex.EncodeToString(hashed[:])
}
