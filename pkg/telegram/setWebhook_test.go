package telegram_test

import (
	"github.com/dozheiny/microsoft-todo-tg-bot/config"
	"github.com/dozheiny/microsoft-todo-tg-bot/pkg/telegram"
	"testing"
)

func TestSetWebhook(t *testing.T) {
	t.Parallel()

	host, err := config.Get("HOST")
	if err != nil {
		t.Errorf("Got error on get config: %s", err.Error())
	}

	token, err := config.Get("TELEGRAM_TOKEN")
	if err != nil {
		t.Errorf("Got error on get config: %s", err.Error())
	}

	if _, err := telegram.SetWebhook(host, token); err != nil {
		t.Errorf("Got error on set webhook: %s", err.Error())
	}

}
