package errors

import "errors"

var ErrInvontoryNotFound = errors.New("inventory not found")
var ErrProductNotFound = errors.New("product not found")
var ErrInventoryQtyNotEnough = errors.New("inventory qty not enough")
