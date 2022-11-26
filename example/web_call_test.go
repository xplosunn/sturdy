package example

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	s "github.com/xplosunn/sturdy/shortsyntax"
	"github.com/xplosunn/sturdy/sturdytest"
	"io"
	"net/http"
	"testing"
)

// TestWebCall checks that the WebCall method is resilient against failure before and after the http call was done
func TestWebCall(t *testing.T) {
	b := sturdytest.ErrorBeforeFirstCallAndThenIncreaseAfterEachCall()
	sturdytest.SturdyTest(t, b, func(ctx context.Context) error {
		callResult, err := WebCall(ctx)
		if err != nil {
			return err
		}
		if callResult == "" {
			return errors.New("empty result")
		}
		return nil
	})
	assert.Empty(t, b.NonInjectedErrors())
}

// WebCall gets a random cat fact from the internet and returns it.
func WebCall(ctx context.Context) (string, error) {
	resp, err := s.S1(ctx, func() (*http.Response, error) {
		return http.Get("https://catfact.ninja/fact")
	})
	if err != nil {
		return "", err
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
