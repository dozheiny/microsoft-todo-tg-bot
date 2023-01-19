package web_test

import (
	"github.com/dozheiny/microsoft-todo-tg-bot/pkg/web"
	"net/http"
	"testing"
)

type TModel struct {
	CurrentPage int `json:"current_page"`
	Data        []struct {
		Breed   string `json:"breed"`
		Country string `json:"country"`
		Origin  string `json:"origin"`
		Coat    string `json:"coat"`
		Pattern string `json:"pattern"`
	} `json:"data"`
	FirstPageUrl string `json:"first_page_url"`
	From         int    `json:"from"`
	LastPage     int    `json:"last_page"`
	LastPageUrl  string `json:"last_page_url"`
	Links        []struct {
		Url    *string `json:"url"`
		Label  string  `json:"label"`
		Active bool    `json:"active"`
	} `json:"links"`
	NextPageUrl string      `json:"next_page_url"`
	Path        string      `json:"path"`
	PerPage     string      `json:"per_page"`
	PrevPageUrl interface{} `json:"prev_page_url"`
	To          int         `json:"to"`
	Total       int         `json:"total"`
}

func TestInit(t *testing.T) {
	t.Parallel()

	web.Init()
}

func TestBase_SetHost(t *testing.T) {
	t.Parallel()

	var host string
	host = "https://example.com"

	base := web.Init().SetHost(host)

	if base.Host != host {
		t.Errorf("Error on set Host; excepted: %s, actual: %s", host, base.Host)
	}
}

func TestBase_SetMethod(t *testing.T) {
	t.Parallel()

	base := web.Init().SetMethod(http.MethodPost)

	if base.Method != http.MethodPost {
		t.Errorf("Error on set Method; excepted: %s, actual: %s", http.MethodPost, base.Method)
	}
}

func TestBase_SetMethodDefault(t *testing.T) {
	t.Parallel()

	base := web.Init()

	if base.Method != http.MethodGet {
		t.Errorf("Error on set Method; excepted: %s, actual: %s", http.MethodGet, base.Method)
	}
}

func TestBase_SetOutput(t *testing.T) {
	t.Parallel()

	m := struct {
		ThisIsTest string
	}{}

	base := web.Init().SetOutput(m)
	if base.Output != m {
		t.Errorf("Error on set Output; excepted: %s, actual: %s", m, base.Output)
	}
}

func TestBase_SetParam(t *testing.T) {
	t.Parallel()

	var key, value string

	key = "foo"
	value = "bar"

	b := web.Init().SetParam(key, value)

	if b.Params[0].Key != key || b.Params[0].Value != value {
		t.Errorf("Error on set params; excepted: %s %s, actual: %s", key, value, b.Params[0])
	}
}

func TestBase_Do(t *testing.T) {
	t.Parallel()

	_, err := web.Init().SetHost("https://catfact.ninja/breeds").
		SetParam("limit", "1").SetMethod(http.MethodGet).SetOutput(&TModel{}).Do()

	if err != nil {
		t.Errorf("Error on Do: %s", err.Error())
	}
}
