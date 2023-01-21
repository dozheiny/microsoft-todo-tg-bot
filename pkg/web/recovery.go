package web

import (
	"fmt"
	"net/http"
)

func Recovery(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			_, err := fmt.Fprintf(w, "Recovered from panic: %v", r)
			if err != nil {
				return
			}
		}
	}()
}
