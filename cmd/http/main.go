package main

import (
	"fmt"
	"github.com/dozheiny/microsoft-todo-tg-bot/config"
	"log"
	"net/http"
)

func main() {

	if err := Init(); err != nil {
		panic(err)
	}

	port, err := config.Get("PORT")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Listening on %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))

}
