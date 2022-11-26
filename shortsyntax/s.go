package s

import (
	"context"
	"errors"
	"github.com/xplosunn/sturdy/behaviour"
)

func S0(ctx context.Context, f func() error) error {
	maybeBehaviour := ctx.Value(behaviour.ContextKey)
	if maybeBehaviour == nil {
		return f()
	}

	b, ok := maybeBehaviour.(behaviour.Behaviour)
	if !ok {
		return errors.New("unexpected in context something that isn't a behaviour for key " + behaviour.ContextKey)
	}

	err := b.BeforeInvocation()
	if err != nil {
		return err
	}
	err = f()
	if err != nil {
		b.NonInjectedError(err)
		return err
	}
	err = b.AfterSuccessfulInvocation()
	if err != nil {
		return err
	}
	return nil
}

func S1[Result1 any](ctx context.Context, f func() (Result1, error)) (Result1, error) {
	maybeBehaviour := ctx.Value(behaviour.ContextKey)
	if maybeBehaviour == nil {
		return f()
	}

	var emptyResult1 Result1

	b, ok := maybeBehaviour.(behaviour.Behaviour)
	if !ok {
		return emptyResult1, errors.New("unexpected in context something that isn't a behaviour for key " + behaviour.ContextKey)
	}

	err := b.BeforeInvocation()
	if err != nil {
		return emptyResult1, err
	}
	result1, err := f()
	if err != nil {
		b.NonInjectedError(err)
		return emptyResult1, err
	}
	err = b.AfterSuccessfulInvocation()
	if err != nil {
		return emptyResult1, err
	}
	return result1, nil
}

func S2[Result1 any, Result2 any](ctx context.Context, f func() (Result1, Result2, error)) (Result1, Result2, error) {
	maybeBehaviour := ctx.Value(behaviour.ContextKey)
	if maybeBehaviour == nil {
		return f()
	}

	var emptyResult1 Result1
	var emptyResult2 Result2

	b, ok := maybeBehaviour.(behaviour.Behaviour)
	if !ok {
		return emptyResult1, emptyResult2, errors.New("unexpected in context something that isn't a behaviour for key " + behaviour.ContextKey)
	}

	err := b.BeforeInvocation()
	if err != nil {
		return emptyResult1, emptyResult2, err
	}
	result1, result2, err := f()
	if err != nil {
		b.NonInjectedError(err)
		return emptyResult1, emptyResult2, err
	}
	err = b.AfterSuccessfulInvocation()
	if err != nil {
		return emptyResult1, emptyResult2, err
	}
	return result1, result2, nil
}

func S3[Result1 any, Result2 any, Result3 any](ctx context.Context, f func() (Result1, Result2, Result3, error)) (Result1, Result2, Result3, error) {
	maybeBehaviour := ctx.Value(behaviour.ContextKey)
	if maybeBehaviour == nil {
		return f()
	}

	var emptyResult1 Result1
	var emptyResult2 Result2
	var emptyResult3 Result3

	b, ok := maybeBehaviour.(behaviour.Behaviour)
	if !ok {
		return emptyResult1, emptyResult2, emptyResult3, errors.New("unexpected in context something that isn't a behaviour for key " + behaviour.ContextKey)
	}

	err := b.BeforeInvocation()
	if err != nil {
		return emptyResult1, emptyResult2, emptyResult3, err
	}
	result1, result2, result3, err := f()
	if err != nil {
		b.NonInjectedError(err)
		return emptyResult1, emptyResult2, emptyResult3, err
	}
	err = b.AfterSuccessfulInvocation()
	if err != nil {
		return emptyResult1, emptyResult2, emptyResult3, err
	}
	return result1, result2, result3, nil
}
