package model

import "fmt"

//ArgumentMismatch error is return whenever a supplied slice is a different
//length than required
type ArgumentMismatch struct {
	weights int
	inputs  int
}

func (w ArgumentMismatch) Error() string {
	return fmt.Sprintf(
		"Numbers of inputs differs from expectation. Expected %d, but received %d.",
		w.weights,
		w.inputs)
}

//NewMismatchError returns a new instance of ArgumentMismatch
func NewMismatchError(numOfValues int, numOfInputs int) ArgumentMismatch {
	return ArgumentMismatch{numOfValues, numOfInputs}
}
