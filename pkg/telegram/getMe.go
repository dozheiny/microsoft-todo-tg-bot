package telegram

import (
	"fmt"
	"github.com/dozheiny/microsoft-todo-tg-bot/models/telegram"
	"github.com/dozheiny/microsoft-todo-tg-bot/pkg/web"
)

// GetMe gets telegram token and returns getMe response from telegram.
func GetMe(id string) (*telegram.GetMe, error) {
	url := fmt.Sprintf("%sbot%s%s", baseUrl, id, getMe)

	model, err := web.Init().SetHost(url).SetOutput(&telegram.GetMe{}).Do()
	if err != nil {
		return nil, err
	}

	output, ok := model.(*telegram.GetMe)
	if !ok {
		return nil, cannotConvertInterfaceToModel
	}

	return output, nil
}
