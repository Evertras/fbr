package sprite

import (
	"image"
	"strings"
	"testing"
)

func TestLoadFramesFromReader(t *testing.T) {
	tests := []struct {
		Name   string
		Input  string
		Error  error
		Output []image.Rectangle
	}{
		{
			Name:  "Simple single frame",
			Input: "10 15 20 30",
			Output: []image.Rectangle{
				image.Rectangle{
					Min: image.Point{
						X: 10,
						Y: 15,
					},
					Max: image.Point{
						X: 30,
						Y: 45,
					},
				},
			},
		},
		{
			Name:   "Empty string",
			Input:  "",
			Output: []image.Rectangle{},
		},
		{
			Name: "Multiple frames with empty final line",
			Input: `10 15 20 30
20 25 20 30
`,
			Output: []image.Rectangle{
				image.Rectangle{
					Min: image.Point{
						X: 10,
						Y: 15,
					},
					Max: image.Point{
						X: 30,
						Y: 45,
					},
				},
				image.Rectangle{
					Min: image.Point{
						X: 20,
						Y: 25,
					},
					Max: image.Point{
						X: 40,
						Y: 55,
					},
				},
			},
		},
	}

	for _, test := range tests {
		reader := strings.NewReader(test.Input)

		frames, err := framesFromReader(reader)

		if test.Error != nil {
			if err == nil {
				t.Errorf("%s: Expected error %q but got nil error", test.Name, test.Error.Error())
			}

			continue
		}

		if err != nil {
			t.Errorf("%s: Got unexpected error: %s", test.Name, err.Error())
			continue
		}

		if len(test.Output) != len(frames) {
			t.Errorf("%s: Expected %d frames but got %d", test.Name, len(test.Output), len(frames))
		}

		for i, r := range frames {
			if r.Min.X != test.Output[i].Min.X {
				t.Errorf("%s: Expected Min.X of %d but got %d", test.Name, test.Output[i].Min.X, r.Min.X)
			}

			if r.Min.Y != test.Output[i].Min.Y {
				t.Errorf("%s: Expected Min.Y of %d but got %d", test.Name, test.Output[i].Min.Y, r.Min.Y)
			}

			if r.Max.X != test.Output[i].Max.X {
				t.Errorf("%s: Expected Max.X of %d but got %d", test.Name, test.Output[i].Max.X, r.Max.X)
			}

			if r.Max.Y != test.Output[i].Max.Y {
				t.Errorf("%s: Expected Max.Y of %d but got %d", test.Name, test.Output[i].Max.Y, r.Max.Y)
			}
		}
	}
}
