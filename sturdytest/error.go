package sturdytest

type TestInjectedError struct {
}

func (err TestInjectedError) Error() string {
	return "(TestInjectedError) this shouldn't happen in production"
}

