package w3utils

import "github.com/pkg/errors"

var (
	ErrUnpackFailed   = errors.New("failed to unpack log data")
	ErrFieldNotFound  = errors.New("field not found")
	ErrNoSuchABIEvent = errors.Wrap(ErrFieldNotFound, "no such ABI event")
)
