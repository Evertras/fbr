// +build js,wasm

package asset

import (
	"io"
	"net/http"

	"github.com/pkg/errors"
)

func readerFromPath(path string) (io.ReadCloser, error) {
	resp, err := http.Get(path)

	if err != nil {
		return nil, errors.Wrap(err, path)
	}

	if resp.StatusCode/100 != 2 {
		return nil, errors.Errorf("Could not load path %q: Status code %d", path, resp.StatusCode)
	}

	return resp.Body, nil
}
