package domain

type ErrorResp struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

type SuccessResp struct {
	ID string `json:"id"`
}
