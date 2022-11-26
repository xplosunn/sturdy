package s_test

import (
	"context"
	"errors"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/assert"
	"github.com/xplosunn/sturdy/behaviour"
	s "github.com/xplosunn/sturdy/shortsyntax"
	"testing"
)

func TestPure(t *testing.T) {
	ctx := context.Background()

	err := s.S0(ctx, func() error {
		return nil
	})
	assert.NoError(t, err)

	expected := ksuid.New().String()
	actual, err := s.S1(ctx, func() (string, error) {
		return expected, nil
	})
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

type BeforeInvocationConstErr struct{}

func (b *BeforeInvocationConstErr) BeforeInvocation() error {
	return errors.New("before invocation err")
}

func (b *BeforeInvocationConstErr) AfterSuccessfulInvocation() error {
	panic("unexpected call to BeforeInvocationConstErr.AfterSuccessfulInvocation")
}

func (b *BeforeInvocationConstErr) NonInjectedError(err error) {
	panic("unexpected call to BeforeInvocationConstErr.NonInjectedError")
}

func TestBeforeInvocationConstErr(t *testing.T) {
	ctx := behaviour.WithBehaviour(context.Background(), &BeforeInvocationConstErr{})

	expected := "before invocation err"
	actual, err := s.S1(ctx, func() (string, error) {
		return expected, nil
	})
	assert.NotNil(t, err)
	assert.Zero(t, actual)
	assert.Equal(t, expected, err.Error())
}

type AfterInvocationConstErr struct {
	NonInjectedErr error
}

func (b *AfterInvocationConstErr) BeforeInvocation() error {
	return nil
}

func (b *AfterInvocationConstErr) AfterSuccessfulInvocation() error {
	return errors.New("after invocation err")
}

func (b *AfterInvocationConstErr) NonInjectedError(err error) {
	b.NonInjectedErr = err
}

func TestAfterInvocationConstErr(t *testing.T) {
	b := AfterInvocationConstErr{}
	ctx := behaviour.WithBehaviour(context.Background(), &b)

	expected := "after invocation err"
	actual, err := s.S1(ctx, func() (string, error) {
		return expected, nil
	})
	assert.NotNil(t, err)
	assert.Zero(t, actual)
	assert.Equal(t, expected, err.Error())
	assert.Nil(t, b.NonInjectedErr)

	expected = "non injected err"
	actual, err = s.S1(ctx, func() (string, error) {
		return "", errors.New("non injected err")
	})
	assert.NotNil(t, err)
	assert.Zero(t, actual)
	assert.Equal(t, expected, err.Error())
	assert.Equal(t, expected, b.NonInjectedErr.Error())
}
