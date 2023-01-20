package telegram_test

import (
	"github.com/dozheiny/microsoft-todo-tg-bot/config"
	"github.com/dozheiny/microsoft-todo-tg-bot/pkg/telegram"
	"testing"
)

func TestGetMe(t *testing.T) {
	t.Parallel()

	id, err := config.Get("TELEGRAM_TOKEN")
	if err != nil {
		t.Errorf("Got Error on get config: %s", err.Error())
	}

	m, err := telegram.GetMe(id)
	if err != nil {
		t.Errorf("Got error on get me function from telegram package: %s", err.Error())
	}

	if !m.Ok {
		t.Errorf("token is invalid.")
	}

}
