package enumerrs

import "errors"

// This package defines custom error types for use in the generated code.

// ErrValueInvalid is returned when a value does not belong to the set of valid values for a type.
var ErrValueInvalid = errors.New("the input value is not valid for the type")
