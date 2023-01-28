package telegram

import (
	"fmt"
	"github.com/dozheiny/microsoft-todo-tg-bot/pkg/web"
	"net/http"
)

func SendMessage(chatId, text, id string) error {

	url := fmt.Sprintf("%sbot%s%s", baseUrl, id, sendMessage)

	_, err := web.Init().
		SetHost(url).
		SetParam("text", text).
		SetParam("chat_id", chatId).
		SetMethod(http.MethodPost).Do()
	if err != nil {
		return err
	}

	return nil
}
