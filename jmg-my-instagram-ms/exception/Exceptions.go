package exception

import "errors"

var ProductAlreadyExists = errors.New("Product already exists")
var ProductNotFound = errors.New("Product not found")
var UserNotFount = errors.New("User not found")
var MultimediaNotFound = errors.New("Multimedia not found")
var LocationNotFound = errors.New("Location not found")
