package tile

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type tileMaker struct {
	tilesURL  string
	outputDir string
	fontsDir  string
}

func NewTileMaker(tilesURL, outputDir, fontsDir string) (*tileMaker, error) {
	info, err := os.Stat(outputDir)
	if err != nil {
		return nil, errors.Wrap(err, "NewTileMaker")
	}

	if !info.IsDir() {
		return nil, errors.New("NewTileMaker: outputDir is not a directory")
	}

	return &tileMaker{
		tilesURL:  tilesURL,
		outputDir: outputDir,
		fontsDir:  fontsDir,
	}, nil
}

func (t *tileMaker) MakeTile(spec Spec) (URl string, err error) {
	filename := fmt.Sprintf("%s.png", spec.ID())
	outputFileName := fmt.Sprintf("%s/%s", t.outputDir, filename)

	if _, err := os.Stat(outputFileName); err != nil {
		img, err := t.createImage(spec)
		if err != nil {
			return "", errors.Wrap(err, "tileMaker.MakeTile")
		}

		if err := img.SavePNG(outputFileName); err != nil {
			return "", errors.Wrap(err, "tileMaker.MakeTile: savePNG")
		}
	}

	url := fmt.Sprintf("%s/%s", t.tilesURL, filename)
	return url, nil
}

func (t *tileMaker) createImage(spec Spec) (*gg.Context, error) {
	var backgroundImage image.Image
	dc := gg.NewContext(1200, 628)
	if spec.BackgroundImage != "" {
		path, err := pathToImage(spec.BackgroundImage)
		if err != nil {
			return nil, errors.Wrap(err, "createImage")
		}

		backgroundImage, err = gg.LoadImage(path)
		if err != nil {
			return nil, errors.Wrap(err, "contextFromImage")
		}

		backgroundImage = imaging.Fill(backgroundImage, dc.Width(), dc.Height(), imaging.Center, imaging.Lanczos)
		dc.DrawImage(backgroundImage, 0, 0)
	} else if spec.BackgroundColor != nil {
		dc.SetColor(spec.BackgroundColor)
		dc.Fill()
	}

	margin := 20.0
	x := margin
	y := margin
	w := float64(dc.Width()) - (2.0 * margin)
	h := float64(dc.Height()) - (2.0 * margin)

	if spec.OverlayColor != nil {
		dc.SetColor(spec.OverlayColor)
		dc.DrawRectangle(x, y, w, h)
		dc.Fill()
	}

	fontPath := filepath.Join(t.fontsDir, "OpenSans-Bold.ttf")
	if err := dc.LoadFontFace(fontPath, 80); err != nil {
		return nil, errors.Wrap(err, "contextFromImage")
	}

	if spec.TextColor == nil {
		spec.TextColor = color.White
	}
	dc.SetColor(spec.TextColor)

	centerX := dc.Width() / 2
	centerY := dc.Height() / 2
	dc.DrawStringWrapped(spec.Text, float64(centerX), float64(centerY), 0.5, 0.5, w, 1.5, gg.AlignCenter)

	return dc, nil
}

func pathToImage(path string) (string, error) {
	if filepath.IsAbs(path) {
		return path, nil
	}

	_, err := url.Parse(path)
	if err != nil {
		return "", errors.New("pathToImage: path is not a valid URL or absolute path")
	}

	path, err = downloadTempFile(path)
	if err != nil {
		return "", errors.Wrap(err, "pathToImage")
	}

	return path, nil
}

func validateImageURL(URL string) error {
	response, err := http.Head(URL)
	if err != nil {
		return errors.Wrap(err, "validateImageURL")
	}

	if response.StatusCode != http.StatusOK {
		return errors.New("Received non 200 response code")
	}

	contentType := response.Header.Get("Content-Type")
	if contentType != "image/png" && contentType != "image/jpeg" {
		return errors.New("URL does not point to a jpg/png image")
	}

	return nil
}

func downloadTempFile(URL string) (path string, err error) {
	if err := validateImageURL(URL); err != nil {
		return "", errors.Wrap(err, "downloadTempFile")
	}

	response, err := http.Get(URL)
	if err != nil {
		return "", errors.Wrap(err, "downloadTempFile")
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", errors.New("Received non 200 response code")
	}

	fileName := fmt.Sprintf("tmp-image-%s", uuid.New().String())
	absolutePath := filepath.Join(os.TempDir(), fileName)
	file, err := os.Create(absolutePath)
	if err != nil {
		return "", errors.Wrap(err, "downloadTempFile")
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return "", errors.Wrap(err, "downloadTempFile")
	}

	return absolutePath, nil
}
