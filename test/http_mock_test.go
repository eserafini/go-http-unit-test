package test

import (
	"context"
	"eserafini/go-http-unit-test/helpers"
	"net/http"
	"testing"
	"time"
)

type PingModel struct {
	Ping string `json:"ping"`
}

type API struct {
	URL string
}

func (api *API) Ping() (*PingModel, error) {
	ping := &PingModel{}

	to := time.Duration(10)
	opt := &helpers.HttpOptions{
		Ctx:    context.Background(),
		Url:    api.URL + "/ping",
		TO:     &to,
		Method: http.MethodGet,
	}

	_, err := helpers.DoRequest(opt, ping)
	return ping, err
}

func TestHttpMockString(t *testing.T) {
	srv := helpers.HttpMock("/ping", http.StatusOK, `{"ping": "pong"}`)
	defer srv.Close()

	api := API{URL: srv.URL}

	res, err := api.Ping()
	if err != nil {
		t.Error(err)
	}

	if err != nil {
		t.Error("expected", nil, "got", err.Error())
	}
	if res.Ping != "pong" {
		t.Error("expected pong got:", res.Ping)
	}
}

func TestHttpMockStruct(t *testing.T) {
	ping := PingModel{"pong"}
	srv := helpers.HttpMock("/ping", http.StatusOK, ping)
	defer srv.Close()

	api := API{URL: srv.URL}

	res, err := api.Ping()
	if err != nil {
		t.Error(err)
	}

	if err != nil {
		t.Error("expected", nil, "got", err.Error())
	}
	if res.Ping != "pong" {
		t.Error("expected pong got:", res.Ping)
	}
}

func TestHttpMockPtrStruct(t *testing.T) {
	ping := &PingModel{"pong"}
	srv := helpers.HttpMock("/ping", http.StatusOK, ping)
	defer srv.Close()

	api := API{URL: srv.URL}

	res, err := api.Ping()
	if err != nil {
		t.Error(err)
	}

	if err != nil {
		t.Error("expected", nil, "got", err.Error())
	}
	if res.Ping != "pong" {
		t.Error("expected pong got:", res.Ping)
	}
}

func TestHttpMockUnknownType(t *testing.T) {
	srv := helpers.HttpMock("/ping", http.StatusOK, 100)
	defer srv.Close()

	api := API{URL: srv.URL}

	res, err := api.Ping()
	if err != nil {
		t.Error(err)
	}

	if err != nil {
		t.Error("expected", nil, "got", err.Error())
	}
	if res.Ping != "" {
		t.Error("expected empty got:", res.Ping)
	}
}
