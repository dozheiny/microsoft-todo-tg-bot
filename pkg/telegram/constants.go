package telegram

import "errors"

const (
	baseUrl = "https://api.telegram.org/"
	getMe   = "/getMe"
)

var (
	cannotConvertInterfaceToModel = errors.New("cannot convert interface to model")
)
