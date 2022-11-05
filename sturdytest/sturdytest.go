package sturdytest

import (
	"context"
	"errors"
	"github.com/xplosunn/sturdy/behaviour"
	"strconv"
	"testing"
)

const maxAttempts = 100

func SturdyTest(
	t *testing.T,
	b behaviour.Behaviour,
	propertyCheck func(t *testing.T),
	testCode func(ctx context.Context) error,
) {
	ctx := context.Background()
	ctx = behaviour.WithBehaviour(ctx, b)
	hadSuccessfulRun := false
	for i := 0; !hadSuccessfulRun && i < maxAttempts; i++ {
		err := testCode(ctx)
		if err == nil {
			hadSuccessfulRun = true
			return
		}
		if !errors.Is(err, TestInjectedError{}) {
			t.Error(err)
			t.FailNow()
		}
		propertyCheck(t)
	}
	if !hadSuccessfulRun {
		t.Error("Didn't finish a single run without injected errors being returned")
	}
}

func SturdyTestSuite(
	t *testing.T,
	b behaviour.Behaviour,
	propertyCheck func(t *testing.T),
	testCode func(ctx context.Context) error,
) {
	ctx := context.Background()
	ctx = behaviour.WithBehaviour(ctx, b)
	hadSuccessfulRun := false
	for i := 0; !hadSuccessfulRun && i < maxAttempts; i++ {
		t.Run("Run "+strconv.Itoa(i+1), func(t *testing.T) {
			err := testCode(ctx)
			if err == nil {
				hadSuccessfulRun = true
				return
			}
			if !errors.Is(err, TestInjectedError{}) {
				t.Error(err)
				t.FailNow()
			}
			propertyCheck(t)
		})
	}
	if !hadSuccessfulRun {
		t.Error("Didn't finish a single run without injected errors being returned")
	}
}
