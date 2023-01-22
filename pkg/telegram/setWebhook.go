package telegram

import (
	"fmt"
	"github.com/dozheiny/microsoft-todo-tg-bot/models/telegram"
	"github.com/dozheiny/microsoft-todo-tg-bot/pkg/web"
)

func SetWebhook(host, token string) (*telegram.SetWebhook, error) {
	url := fmt.Sprintf("%sbot%s%s", baseUrl, token, setWebhook)

	o, err := web.Init().SetHost(url).SetHeader("X-Telegram-Bot-Api-Secret-Token", token).
		SetParam("url", host).SetOutput(&telegram.SetWebhook{}).Do()
	if err != nil {
		return nil, err
	}

	output, ok := o.(*telegram.SetWebhook)
	if !ok {
		return nil, cannotConvertInterfaceToModel
	}

	if !output.Ok {
		return output, cannotSetWebhook
	}

	return output, nil
}
