package transform

type TransformError struct {
	message string
}

var _ error = TransformError{}

func (t TransformError) Error() string {
	return t.message
}

func newError(msg string) TransformError {
	return TransformError{message: msg}
}

var InsufficientDataError = newError("mino: insufficient data to perform this calculation.")
