package telegram

import (
	"encoding/json"
	"fmt"
	"github.com/dozheiny/microsoft-todo-tg-bot/config"
	"github.com/dozheiny/microsoft-todo-tg-bot/models/telegram"
	telegram2 "github.com/dozheiny/microsoft-todo-tg-bot/pkg/telegram"
	"github.com/dozheiny/microsoft-todo-tg-bot/pkg/web"
	"io"
	"log"
	"net/http"
	"strconv"
)

// Update will receive all requests from https://api.telegram.org
func Update(w http.ResponseWriter, r *http.Request) {
	defer web.Recovery(w, r)

	// read body and parse it into json model.
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Print(err.Error())

		_, err := io.WriteString(w, fmt.Sprintf("Got error on read body: %s", err.Error()))
		if err != nil {
			log.Print(fmt.Sprintf("got error on io.WriteString: %s", err.Error()))
		}
	}

	inputForm := new(telegram.Update)

	if err := json.Unmarshal(body, inputForm); err != nil {
		log.Print(fmt.Sprintf("got error on parse json un marshal: %s", err.Error()))

		_, err := io.WriteString(w, fmt.Sprintf("Got error on parse body: %s", err.Error()))
		if err != nil {
			log.Print(fmt.Sprintf("got error on io.WriteString: %s", err.Error()))
		}
	}

	token, err := config.Get("TELEGRAM_TOKEN")
	if err != nil {
		log.Print(fmt.Sprintf("got error on reading telegram token :%s", err.Error()))
	}

	text := fmt.Sprintf("Hi %s", inputForm.Message.From.FirstName)

	if err := telegram2.SendMessage(strconv.
		FormatInt(inputForm.Message.Chat.ID, 10), text, token); err != nil {
		log.Print(err)
	}

}
