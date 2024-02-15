# frameserver-go
A Farcaster frame server implementation written in Go.

----------------

This repository is meant to demonstrate how to implement a Farcaster frame server with Golang. There are two main 
components:

1. A server that produces [farcaster-compliant](https://docs.farcaster.xyz/reference/frames/spec) frame markup.
2. A "tile maker" utility for creating simple image tiles. 

This is a "hello world" demonstration. You will find the endpoints served by the http server inside the `greeting` package.

The main package and application entrypoint are in `cmd/frameserver/main.go`.

### Creating an image tile:
```go
tileMaker, err := tile.NewTileMaker(tilesURL, outputDir, fontsDir)
if err != nil {
	log.Fatal(err)
}

tileSpec := tile.Spec{
    Text:            "gm farcaster!",
    BackgroundImage: "/path/to/background.jpg", // can be a URL to a jpg/png 
    TextColor:       color.White,
    OverlayColor:    color.RGBA{R: 0, G: 0, B: 0, A: 80},
}

tileURL, err := tileMaker.MakeTile(tileSpec)
```

### Creating a frame:
```go
frame := farcaster.Frame{
    Version: "vNext",
	PostURL: "https://example.com/post",
    Image: farcaster.Image{
        URL:         "https://example.com/static/tiles/abc123.png",
        AspectRatio: farcaster.AspectRatio_2_1,
    },
    Buttons: []farcaster.Button{
        farcaster.NewPostButton("reveal"),
    },
}
```

### Running the demo server:
```bash
docker build -t frameserver-go . && \
docker run \
  -e PORT=8080 \
  -e APP_URL=http://127.0.0.1:8080 \
  -e HUB_URL=https://nemes.farcaster.xyz:2281 \
  -e STATIC_DIR=/root/static \
  -p 8080:8080 \
  frameserver-go
  
```
