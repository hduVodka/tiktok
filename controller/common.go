package controller

type Resp struct {
	StatusCode int    `json:"statusCode"`
	StatusMsg  string `json:"statusMsg,omitempty"`
}

var ErrInvalidParams = "invalid params"
