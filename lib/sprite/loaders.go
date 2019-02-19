package sprite

import (
	"image"
	"net/http"

	"github.com/hajimehoshi/ebiten"
	"github.com/pkg/errors"
)

// StaticFromPath loads a static sprite from the given relative path
func StaticFromPath(path string) (Sprite, error) {
	// TODO: This could be a file system or something else based on build
	// environment (mobile!), but for now it's just an HTTP request
	resp, err := http.Get(path)

	if err != nil {
		return nil, errors.Wrap(err, path)
	}

	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		return nil, errors.Errorf("Could not load path %q: Status code %d", path, resp.StatusCode)
	}

	raw, _, err := image.Decode(resp.Body)

	img, err := ebiten.NewImageFromImage(raw, ebiten.FilterDefault)

	return NewStatic(img), nil
}
