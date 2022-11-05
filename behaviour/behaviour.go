package behaviour

import "context"

const ContextKey = "TEST_STURDY_BEHAVIOUR"

type Behaviour interface {
	BeforeInvocation() error
	AfterSuccessfulInvocation() error
	//TODO think about this
	NonInjectedError(err error)
}

func WithBehaviour(ctx context.Context, b Behaviour) context.Context {
	return context.WithValue(ctx, ContextKey, b)
}
