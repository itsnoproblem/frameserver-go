package farcaster

type Image struct {
	URL         string
	AspectRatio AspectRatio
}

type AspectRatio string

func (s AspectRatio) String() string {
	return string(s)
}
