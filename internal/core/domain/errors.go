package domain

import "errors"

var (
	ErrInternalError = errors.New("internal server error")
)

var (
	ErrorMessageFailedToConvertToJSON     = "failed to convert to json"
	ErrorMessageFailedToConvertToGoStruct = "failed to convert to go struct"
	ErrorMessageFailedToInsertData        = "failed to insert payload to db"
	ErrorMessageFailedToUpdateData        = "failed to update payload to db"
)
