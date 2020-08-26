package constants

import (
	"errors"

	errors2 "github.com/pkg/errors"
)

var (
	ErrAccessDenied       = errors.New("access denied")
	ErrAttemptsIsExceeded = errors2.Wrap(ErrAccessDenied, "attempts is exceeded")
)
