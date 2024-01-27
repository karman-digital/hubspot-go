package hsvalidate

import "errors"

var ErrMismatchedSignatures = errors.New("signatures mismatched")
var ErrTimestampTooOld = errors.New("timestamp more than 5 minutes old")
var ErrTimestampInvalid = errors.New("invalid timestamp")
