package client

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

func NewWSClient(entrypoint, username, password string) (*websocket.Conn, error) {
	authString := strings.Join([]string{username, password}, ":")
	credentials := base64.StdEncoding.EncodeToString([]byte(authString))

	c, _, err := websocket.DefaultDialer.Dial(entrypoint, http.Header{
		"User-Agent":    []string{"socket_cat/1.0"},
		"Authorization": []string{"Basic " + credentials},
	})
	if err != nil {
		return nil, err
	}

	return c, nil
}
