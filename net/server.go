package net

import "fmt"

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
}

func NewServer(name string) *Server {
	return &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
}

func (s *Server) Start() {
	fmt.Printf("[START] Server Listenner at IP: %s, Port: %d, is starting\n", s.IP, s.Port)
}
