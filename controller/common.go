package controller

type Resp struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

var ErrInvalidParams = "invalid params"
var ErrUserAlreadyExist = "user already exist"
var ErrInsertFailed = "insert failed"
var ErrIncorrectPassword = "incorrect password"
var ErrUserNotFound = "user not found"
var ErrFormatError = "format error"
var ErrInternalServer = "internal server error"
