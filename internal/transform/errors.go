package transform

import (
	"fmt"
)

type TransformerError struct {
	// Transformer that returned the error
	Transformer Transformer

	// Err is the actual error
	Err error

	// Input is the input string to be transformed
	Input string
}

func (err *TransformerError) Error() string {
	return fmt.Sprintf("Transformer <%T> failed on input '%s': %s",
		err.Transformer,
		err.Input,
		err.Err)
}
