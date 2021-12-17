package server

import "net/http"

type Server struct {
	Host string
	Port string
}

func (s *Server) Start(router http.Handler) error {
	return http.ListenAndServe(s.Host+":"+s.Port, router)
}
