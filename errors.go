package iter

import "errors"

// Error to indicate the iterator is at the end of it's sequence.
var ErrAtEnd = errors.New("iterator is at the end")

// Error to indicate that the bound iterator is attempting to have an element
// accessed that is out of it's bounds.
var ErrNotInBounds = errors.New("index is outside of iterator bounds")

// Error to indicate the value passed into the Range iterator is invalid.
var ErrInvalidRangeValue = errors.New("range iterator can not be constructed with values less than 0")
