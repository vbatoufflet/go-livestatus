package livestatus

import (
	"errors"
)

// Record retrieval errors
var (
	ErrUnknownColumn = errors.New("unknown record column")
	ErrInvalidQuery  = errors.New("invalid query")
	ErrInvalidValue  = errors.New("invalid record value")
)
