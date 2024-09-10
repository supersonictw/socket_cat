package server

import (
	"encoding/json"
	"net/http"

	"github.com/olahol/melody"
)

type CatMux struct {
	wss *melody.Melody
}

func (c *CatMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/nodes" {
		c.nodes(w, r)
		return
	}
	if r.URL.Path == "/exchange" {
		c.exchange(w, r)
		return
	}
	http.NotFound(w, r)
}

func (c *CatMux) nodes(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{})
}

func (c *CatMux) exchange(w http.ResponseWriter, r *http.Request) {
	c.wss.HandleRequest(w, r)
}

func NewServer(addr string) {
	m := melody.New()
	m.HandleConnect(func(s *melody.Session) {
		src, _, _ := s.Request.BasicAuth()
		dst := s.Request.URL.Query().Get("dst")

		s.Set("srcName", src)
		s.Set("dstName", dst)

		if dst == "" {
			return
		}

		sessions, err := m.Sessions()
		if err != nil {
			panic(err)
		}

		for _, ds := range sessions {
			sSrc, sExists := ds.Get("srcName")
			if !sExists || sSrc != dst {
				continue
			}
			ds.Set("dstName", src)
			ds.Set("dstSession", s)
			s.Set("dstSession", ds)
		}
	})
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		ds, dsExists := s.Get("dstSession")
		if !dsExists {
			panic("")
		}

		session, dsOk := ds.(*melody.Session)
		if !dsOk {
			panic("")
		}

		session.Write(msg)
	})
	m.HandleMessageBinary(func(s *melody.Session, b []byte) {
		ds, dsExists := s.Get("dstSession")
		if !dsExists {
			panic("")
		}

		session, dsOk := ds.(*melody.Session)
		if !dsOk {
			panic("")
		}

		session.WriteBinary(b)
	})

	c := &CatMux{m}
	http.ListenAndServe(addr, c)
}
