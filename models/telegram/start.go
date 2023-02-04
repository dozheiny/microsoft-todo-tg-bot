package telegram

import (
	"fmt"
	"github.com/dozheiny/microsoft-todo-tg-bot/config"
	telegram2 "github.com/dozheiny/microsoft-todo-tg-bot/pkg/telegram"
	"log"
	"strconv"
)

func (u *Update) Start() {
	text := fmt.Sprintf(`
	Hi %s! this is Microsoft Todo Telegram Bot ✨
	you can use /authorize for authorize your microsoft Todo.
	Have fun 😁
`, u.Message.From.FirstName)

	token, err := config.Get("TELEGRAM_TOKEN")
	if err != nil {
		log.Print(fmt.Sprintf("got error on reading telegram token :%s", err.Error()))
	}

	if err := telegram2.SendMessage(strconv.
		FormatInt(u.Message.Chat.ID, 10), text, token); err != nil {
		log.Print(err)
	}

}
