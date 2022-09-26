package testing

import (
	"net/http"
	"net/http/httptest"
)

func NewScoreRequest(name string) *http.Request {
	return httptest.NewRequest(http.MethodGet, "/players/"+name, nil)
}

func NewPostWinRequest(name string) *http.Request {
	return httptest.NewRequest(http.MethodPost, "/players/Pepper", nil)
}

func NewLeagueRequest() *http.Request {
	return httptest.NewRequest(http.MethodGet, "/league", nil)
}

func NewGameRequest() *http.Request {
	return httptest.NewRequest(http.MethodGet, "/game", nil)
}
