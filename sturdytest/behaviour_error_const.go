package sturdytest

import "github.com/xplosunn/sturdy/behaviour"

func ErrorConstBeforeCall() behaviour.Behaviour {
	return errorConst{
		BeforeCall: true,
	}
}

func ErrorConstAfterCall() behaviour.Behaviour {
	return errorConst{
		BeforeCall: false,
	}
}

type errorConst struct {
	BeforeCall      bool
	NonInjectedErrs []error
}

func (errC errorConst) BeforeInvocation() error {
	if errC.BeforeCall {
		return TestInjectedError{}
	}
	return nil
}

func (errC errorConst) AfterSuccessfulInvocation() error {
	if !errC.BeforeCall {
		return TestInjectedError{}
	}
	return nil
}

func (errC errorConst) NonInjectedError(err error) {
	errC.NonInjectedErrs = append(errC.NonInjectedErrs, err)
}

func (errC errorConst) NonInjectedErrors() []error {
	return errC.NonInjectedErrs
}
