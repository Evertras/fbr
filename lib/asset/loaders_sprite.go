package asset

import (
	"image"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten"
	"github.com/pkg/errors"
)

var imageCache = make(map[string]*ebiten.Image)
var frameCache = make(map[string][]image.Rectangle)

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

// LoadImageFromPath loads a static sprite from the given relative path, caching results for future calls
func LoadImageFromPath(path string) (*ebiten.Image, error) {
	if img, ok := imageCache[path]; ok {
		return img, nil
	}

	reader, err := readerFromPath(path)

	if err != nil {
		return nil, err
	}

	raw, _, err := image.Decode(reader)

	reader.Close()

	if err != nil {
		return nil, err
	}

	img, err := ebiten.NewImageFromImage(raw, ebiten.FilterDefault)

	if err == nil {
		imageCache[path] = img
	}

	return img, err
}

// LoadFramesFromPath loads frame data from the given path, caching results for future calls
func LoadFramesFromPath(path string) ([]image.Rectangle, error) {
	if frames, ok := frameCache[path]; ok {
		return frames, nil
	}

	reader, err := readerFromPath(path)

	if err != nil {
		return nil, err
	}

	frames, err := framesFromReader(reader)

	reader.Close()

	if err == nil {
		frameCache[path] = frames
	}

	return frames, err
}
