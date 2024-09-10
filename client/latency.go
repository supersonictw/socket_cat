package client

import (
	"net/http"
	"time"
)

type HttpLatency struct {
	Endpoint string
	Duration time.Duration
}

func BestLatencyServer(endpoints []string) *HttpLatency {
	client := &http.Client{}

	ch := make(chan *HttpLatency, 1)
	for _, e := range endpoints {
		go func(endpoint string) {
			timeStart := time.Now()
			resp, err := client.Head(endpoint + "/meow")
			if err != nil {
				return
			}
			if resp.StatusCode != http.StatusAccepted {
				return
			}
			timeEnd := time.Now()
			ch <- &HttpLatency{
				Endpoint: endpoint,
				Duration: timeEnd.Sub(timeStart),
			}
		}(e)
	}

	go func() {
		time.Sleep(5 * time.Second)
		ch <- nil
	}()

	defer func() {
		close(ch)
	}()

	return <-ch
}
