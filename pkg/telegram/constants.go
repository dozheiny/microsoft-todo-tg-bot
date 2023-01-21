package telegram

import "errors"

const (
	baseUrl    = "https://api.telegram.org/"
	getMe      = "/getMe"
	setWebhook = "/setWebhook"
)

var (
	cannotConvertInterfaceToModel = errors.New("cannot convert interface to model")
	cannotSetWebhook              = errors.New("cannot set webhook")
)
