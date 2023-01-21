package telegram

import (
	"encoding/json"
	"fmt"
	"github.com/dozheiny/microsoft-todo-tg-bot/models/telegram"
	"github.com/dozheiny/microsoft-todo-tg-bot/pkg/web"
	"io"
	"log"
	"net/http"
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
			panic(err)
		}
	}

	inputForm := new(telegram.Update)

	if err := json.Unmarshal(body, inputForm); err != nil {
		log.Print(err.Error())

		_, err := io.WriteString(w, fmt.Sprintf("Got error on parse body: %s", err.Error()))
		if err != nil {
			panic(err)
		}
	}

	fmt.Println(inputForm)
}
