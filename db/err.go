package db

import "errors"

var ErrDatabase = errors.New("database error")
var ErrUserNotFound = errors.New("user not found")
