package main

import (
	"github.com/dozheiny/microsoft-todo-tg-bot/config"
	"github.com/dozheiny/microsoft-todo-tg-bot/pkg/telegram"
)

// Init will initialize and check telegram connection, microsoft connection
// database connection etc.
func Init() error {

	tgToken, err := config.Get("TELEGRAM_TOKEN")
	if err != nil {
		return err
	}

	getMe, err := telegram.GetMe(tgToken)
	if err != nil {
		return err
	}

	if !getMe.Ok {
		return okFromGetMeShouldBeTrue
	}

	return nil
}
