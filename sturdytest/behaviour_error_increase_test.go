package sturdytest_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/xplosunn/sturdy/behaviour"
	s "github.com/xplosunn/sturdy/shortsyntax"
	"github.com/xplosunn/sturdy/sturdytest"
	"testing"
)

func TestErrorIncreaseBeforeCall(t *testing.T) {
	b := sturdytest.ErrorIncreaseBeforeCall()
	ctx := behaviour.WithBehaviour(context.Background(), b)

	// FailAt:       1
	// FailureCount: 1
	called := false
	err := s.S0(ctx, func() error {
		called = true
		return errors.New("unexpected")
	})
	assert.NotNil(t, err)
	assert.False(t, called)
	assert.Equal(t, "(TestInjectedError) this shouldn't happen in production", err.Error())

	// FailAt:       2
	// FailureCount: 1
	called = false
	err = s.S0(ctx, func() error {
		called = true
		return errors.New("expected")
	})
	assert.NotNil(t, err)
	assert.True(t, called)
	assert.Equal(t, "expected", err.Error())

	// FailAt:       2
	// FailureCount: 2
	called = false
	err = s.S0(ctx, func() error {
		called = true
		return errors.New("unexpected")
	})
	assert.NotNil(t, err)
	assert.False(t, called)
	assert.Equal(t, "(TestInjectedError) this shouldn't happen in production", err.Error())

	// FailAt:       3
	// FailureCount: 1
	called = false
	err = s.S0(ctx, func() error {
		called = true
		return errors.New("expected")
	})
	assert.NotNil(t, err)
	assert.True(t, called)
	assert.Equal(t, "expected", err.Error())

	// FailAt:       3
	// FailureCount: 2
	called = false
	err = s.S0(ctx, func() error {
		called = true
		return errors.New("expected")
	})
	assert.NotNil(t, err)
	assert.True(t, called)
	assert.Equal(t, "expected", err.Error())

	// FailAt:       3
	// FailureCount: 3
	called = false
	err = s.S0(ctx, func() error {
		called = true
		return errors.New("unexpected")
	})
	assert.NotNil(t, err)
	assert.False(t, called)
	assert.Equal(t, "(TestInjectedError) this shouldn't happen in production", err.Error())
}

func TestErrorIncreaseAfterCall(t *testing.T) {
	b := sturdytest.ErrorIncreaseAfterCall()
	ctx := behaviour.WithBehaviour(context.Background(), b)

	// FailAt:       1
	// FailureCount: 1
	called := false
	err := s.S0(ctx, func() error {
		called = true
		return nil
	})
	assert.NotNil(t, err)
	assert.True(t, called)
	assert.Equal(t, "(TestInjectedError) this shouldn't happen in production", err.Error())

	// FailAt:       2
	// FailureCount: 1
	called = false
	err = s.S0(ctx, func() error {
		called = true
		return nil
	})
	assert.Nil(t, err)
	assert.True(t, called)

	// FailAt:       2
	// FailureCount: 2
	called = false
	err = s.S0(ctx, func() error {
		called = true
		return nil
	})
	assert.NotNil(t, err)
	assert.True(t, called)
	assert.Equal(t, "(TestInjectedError) this shouldn't happen in production", err.Error())

	// FailAt:       3
	// FailureCount: 1
	called = false
	err = s.S0(ctx, func() error {
		called = true
		return nil
	})
	assert.Nil(t, err)
	assert.True(t, called)

	// FailAt:       3
	// FailureCount: 2
	called = false
	err = s.S0(ctx, func() error {
		called = true
		return nil
	})
	assert.Nil(t, err)
	assert.True(t, called)

	// FailAt:       3
	// FailureCount: 3
	called = false
	err = s.S0(ctx, func() error {
		called = true
		return nil
	})
	assert.NotNil(t, err)
	assert.True(t, called)
	assert.Equal(t, "(TestInjectedError) this shouldn't happen in production", err.Error())
}

func TestErrorBeforeFirstCallAndThenIncreaseAfterEachCall(t *testing.T) {
	b := sturdytest.ErrorBeforeFirstCallAndThenIncreaseAfterEachCall()
	ctx := behaviour.WithBehaviour(context.Background(), b)

	// HasErroredBeforeFirstCall: false
	called := false
	err := s.S0(ctx, func() error {
		called = true
		return nil
	})
	assert.NotNil(t, err)
	assert.False(t, called)
	assert.Equal(t, "(TestInjectedError) this shouldn't happen in production", err.Error())

	// FailAt:       1
	// FailureCount: 1
	called = false
	err = s.S0(ctx, func() error {
		called = true
		return nil
	})
	assert.NotNil(t, err)
	assert.True(t, called)
	assert.Equal(t, "(TestInjectedError) this shouldn't happen in production", err.Error())

	// FailAt:       2
	// FailureCount: 1
	called = false
	err = s.S0(ctx, func() error {
		called = true
		return nil
	})
	assert.Nil(t, err)
	assert.True(t, called)

	// FailAt:       2
	// FailureCount: 2
	called = false
	err = s.S0(ctx, func() error {
		called = true
		return nil
	})
	assert.NotNil(t, err)
	assert.True(t, called)
	assert.Equal(t, "(TestInjectedError) this shouldn't happen in production", err.Error())

	// FailAt:       3
	// FailureCount: 1
	called = false
	err = s.S0(ctx, func() error {
		called = true
		return nil
	})
	assert.Nil(t, err)
	assert.True(t, called)

	// FailAt:       3
	// FailureCount: 2
	called = false
	err = s.S0(ctx, func() error {
		called = true
		return nil
	})
	assert.Nil(t, err)
	assert.True(t, called)

	// FailAt:       3
	// FailureCount: 3
	called = false
	err = s.S0(ctx, func() error {
		called = true
		return nil
	})
	assert.NotNil(t, err)
	assert.True(t, called)
	assert.Equal(t, "(TestInjectedError) this shouldn't happen in production", err.Error())
}
