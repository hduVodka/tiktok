package controller

type Resp struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

var ErrInvalidParams = "invalid params"
var ErrUserAlreadyExist = "user already exist"
var ErrInsertFailed = "insert failed"
var ErrIcorrectPassword = "incorrect password"
var ErrFormatError = "format error"
