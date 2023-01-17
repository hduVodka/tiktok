package controller

type Resp struct {
	StatusCode int    `json:"statusCode"`
	StatusMsg  string `json:"statusMsg,omitempty"`
}

var ErrInvalidParams = "invalid params"
var ErrUserAlreadyExist = "user already exist"
var ErrInsertFailed = "insert failed"
var ErrIcorrectPassword = "incorrect password"
var ErrFormatError = "format error"
