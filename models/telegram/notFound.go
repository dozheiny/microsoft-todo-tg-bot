package telegram

import (
	"fmt"
	"github.com/dozheiny/microsoft-todo-tg-bot/config"
	telegram2 "github.com/dozheiny/microsoft-todo-tg-bot/pkg/telegram"
	"log"
	"strconv"
)

func (u *Update) NotFound() {
	text := "command not found!"

	token, err := config.Get("TELEGRAM_TOKEN")
	if err != nil {
		log.Print(fmt.Sprintf("got error on reading telegram token :%s", err.Error()))
	}

	if err := telegram2.SendMessage(strconv.
		FormatInt(u.Message.Chat.ID, 10), text, token); err != nil {
		log.Print(err)
	}

}
