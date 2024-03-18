package errors

import "errors"

var ErrCannotPassStatusNotPending = errors.New("cannot pass status not pending")
var ErrCannotRejectStatusNotPending = errors.New("cannot reject status not pending")
var ErrProductNotFound = errors.New("product not found")
