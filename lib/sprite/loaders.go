package sprite

import (
	"image"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/hajimehoshi/ebiten"
	"github.com/pkg/errors"
)

func readerFromPath(path string) (io.ReadCloser, error) {
	// TODO: This could be a file system or something else based on build
	// environment (mobile!), but for now it's just an HTTP request
	resp, err := http.Get(path)

	if err != nil {
		return nil, errors.Wrap(err, path)
	}

	if resp.StatusCode/100 != 2 {
		return nil, errors.Errorf("Could not load path %q: Status code %d", path, resp.StatusCode)
	}

	return resp.Body, nil
}

func imageFromPath(path string) (*ebiten.Image, error) {
	reader, err := readerFromPath(path)

	if err != nil {
		return nil, err
	}

	raw, _, err := image.Decode(reader)

	reader.Close()

	return ebiten.NewImageFromImage(raw, ebiten.FilterDefault)
}

func framesFromReader(reader io.Reader) ([]image.Rectangle, error) {
	raw, err := ioutil.ReadAll(reader)

	if err != nil {
		return nil, errors.Wrap(err, "failed to read from reader:")
	}

	lines := strings.Split(string(raw), "\n")
	rects := make([]image.Rectangle, len(lines))[:0]

	for _, l := range lines {
		if l == "" {
			continue
		}

		fields := strings.Split(l, " ")

		if len(fields) > 1 && len(fields) != 4 {
			return nil, errors.Errorf("unexpected field count of %d", len(fields))
		}

		x, err := strconv.Atoi(fields[0])

		if err != nil {
			return nil, errors.Errorf("unexpected x coordinate %q", fields[0])
		}

		y, err := strconv.Atoi(fields[1])

		if err != nil {
			return nil, errors.Errorf("unexpected y coordinate %q", fields[1])
		}

		rect := image.Rectangle{}

		rect.Min.X = x
		rect.Min.Y = y

		width, err := strconv.Atoi(fields[2])

		if err != nil {
			return nil, errors.Errorf("unexpected width %q", fields[2])
		}

		height, err := strconv.Atoi(fields[3])

		if err != nil {
			return nil, errors.Errorf("unexpected height %q", fields[3])
		}

		rect.Max.X = x + width
		rect.Max.Y = y + height

		rects = append(rects, rect)
	}

	return rects, nil
}

func framesFromPath(path string) ([]image.Rectangle, error) {
	reader, err := readerFromPath(path)

	if err != nil {
		return nil, err
	}

	frames, err := framesFromReader(reader)

	reader.Close()

	return frames, err
}

// StaticFromPath loads a static sprite from the given relative path
func StaticFromPath(path string) (Sprite, error) {
	img, err := imageFromPath(path)

	if err != nil {
		return nil, err
	}

	return NewStatic(img), nil
}

// AnimatedFromPath loads an animated sprite with a given sprite sheet and animation data
func AnimatedFromPath(pathSheet, pathFrames string, opts AnimationOptions) (Sprite, error) {
	var sheet *ebiten.Image
	var frames []image.Rectangle

	// We'll just take the first error and return that
	errs := make(chan error, 1)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		var err error
		sheet, err = imageFromPath(pathSheet)

		if err != nil {
			select {
			case errs <- err:
			default:
			}
		}
	}()

	go func() {
		defer wg.Done()

		var err error
		frames, err = framesFromPath(pathFrames)

		if err != nil {
			select {
			case errs <- err:
			default:
			}
		}
	}()

	wg.Wait()

	select {
	case err := <-errs:
		return nil, err
	default:
	}

	return NewAnimated(sheet, frames, opts), nil
}
