package livestatus

import (
	"errors"
)

// Record retrieval errors
var (
	ErrUnknownColumn = errors.New("unknown record column")
	ErrInvalidValue  = errors.New("invalid record value")
)
