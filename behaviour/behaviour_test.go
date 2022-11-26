package behaviour_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/xplosunn/sturdy/behaviour"
	"testing"
)

func TestWithBehaviour(t *testing.T) {
	ctx := context.Background()
	ctx = behaviour.WithBehaviour(ctx, SampleBehaviour{})
	maybeB := ctx.Value("TEST_STURDY_BEHAVIOUR")
	b, ok := maybeB.(SampleBehaviour)
	assert.True(t, ok)
	assert.Equal(t, SampleBehaviour{}, b)
}

type SampleBehaviour struct{}

func (s SampleBehaviour) BeforeInvocation() error {
	panic("implement me")
}

func (s SampleBehaviour) AfterSuccessfulInvocation() error {
	panic("implement me")
}

func (s SampleBehaviour) NonInjectedError(err error) {
	panic("implement me")
}

func (s SampleBehaviour) NonInjectedErrors() []error {
	panic("implement me")
}
