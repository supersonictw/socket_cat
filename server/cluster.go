package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"sync"
)

type MessageType int

const (
	MessageTypeNodeJoin = 1000 + iota
)

type Message struct {
	Type    MessageType
	Message []byte
}

type Cluster struct {
	client      *http.Client
	entrypoints []string
}

func (c *Cluster) Send(entrypoint string, msg *Message) error {
	jsonBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	_, err = c.client.Post(
		entrypoint+"/messages",
		"application/json",
		bytes.NewReader(jsonBytes),
	)
	return err
}

func (c *Cluster) Broadcast(msg *Message) {
	wg := &sync.WaitGroup{}
	wg.Add(len(c.entrypoints))

	for _, e := range c.entrypoints {
		go func(entrypoint string) {
			defer wg.Done()
			c.Send(entrypoint, msg)
		}(e)
	}

	wg.Wait()
}
