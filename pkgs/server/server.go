package server

import (
	"net/http"
	"sync"
)

type Server struct {
	Host string
	Port string
	Wg   *sync.WaitGroup
}

func NewServer(host, port string) *Server {
	return &Server{
		Host: host,
		Port: port,
		Wg:   &sync.WaitGroup{},
	}
}

func (s *Server) Start(router http.Handler) error {
	return http.ListenAndServe(s.Host+":"+s.Port, router)
}

// Background starts a go routine that runs on the servers pool. In the event of a shutdown
// request, the server will wait until all open goroutines have finished before shutting down.
func (s *Server) Background(task func()) {
	s.Wg.Add(1)
	go func() {
		defer s.Wg.Done()
		task()
	}()
}
