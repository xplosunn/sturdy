package example

import (
	"context"
	"errors"
	s "github.com/xplosunn/sturdy/shortsyntax"
	"github.com/xplosunn/sturdy/sturdytest"
	"io"
	"net/http"
	"testing"
)

func TestWebCall(t *testing.T) {
	sturdytest.SturdyTest(t, sturdytest.ErrorBeforeFirstCallAndThenIncreaseAfterEachCall(), func(ctx context.Context) error {
		callResult, err := WebCall(ctx)
		if err != nil {
			return err
		}
		if callResult == "" {
			return errors.New("empty result")
		}
		return nil
	})
}

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
