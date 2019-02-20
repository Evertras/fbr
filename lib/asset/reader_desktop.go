// +build windows linux darwin

package asset

import (
	"io"
	"os"

	"github.com/pkg/errors"
)

func readerFromPath(path string) (io.ReadCloser, error) {
	file, err := os.Open("front/" + path)

	if err != nil {
		return nil, errors.Wrap(err, path)
	}

	return file, nil
}
