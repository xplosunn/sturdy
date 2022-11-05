package sturdytest

import "github.com/xplosunn/sturdy/behaviour"

func ErrorIncreaseBeforeCall() behaviour.Behaviour {
	return &increasinglyErroring{
		BeforeCall:   true,
		FailAt:       1,
		FailureCount: 0,
	}
}

func ErrorIncreaseAfterCall() behaviour.Behaviour {
	return &increasinglyErroring{
		BeforeCall:   false,
		FailAt:       1,
		FailureCount: 0,
	}
}

type increasinglyErroring struct {
	BeforeCall   bool
	FailAt       int
	FailureCount int
}

func (i *increasinglyErroring) BeforeInvocation() error {
	if i.BeforeCall {
		i.FailureCount += 1
		if i.FailureCount == i.FailAt {
			i.FailAt += 1
			i.FailureCount = 0
			return TestInjectedError{}
		}
	}
	return nil
}

func (i *increasinglyErroring) AfterSuccessfulInvocation() error {
	if !i.BeforeCall {
		i.FailureCount += 1
		if i.FailureCount == i.FailAt {
			i.FailAt += 1
			i.FailureCount = 0
			return TestInjectedError{}
		}
	}
	return nil
}

func (i *increasinglyErroring) NonInjectedError(err error) {

}

func ErrorBeforeFirstCallAndThenIncreaseAfterEachCall() behaviour.Behaviour {
	return &beforeFirstCallAndThenIncreaseAfterEachCallErroring{
		HasErroredBeforeFirstCall: false,
		FailAt:                    1,
		FailureCount:              0,
	}
}

type beforeFirstCallAndThenIncreaseAfterEachCallErroring struct {
	HasErroredBeforeFirstCall bool
	FailAt                    int
	FailureCount              int
}

func (b *beforeFirstCallAndThenIncreaseAfterEachCallErroring) BeforeInvocation() error {
	if !b.HasErroredBeforeFirstCall {
		b.HasErroredBeforeFirstCall = true
		return TestInjectedError{}
	}
	return nil
}

func (b *beforeFirstCallAndThenIncreaseAfterEachCallErroring) AfterSuccessfulInvocation() error {
	b.FailureCount += 1
	if b.FailureCount == b.FailAt {
		b.FailAt += 1
		b.FailureCount = 0
		return TestInjectedError{}
	}
	return nil
}

func (b *beforeFirstCallAndThenIncreaseAfterEachCallErroring) NonInjectedError(err error) {

}
