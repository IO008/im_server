package net

import (
	"fmt"
	"im_server/iface"
	"net"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
	Router    iface.IRouter
}

func NewServer(name string) *Server {
	return &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
		Router:    nil,
	}
}

func (s *Server) Start() {
	fmt.Printf("[START] Server Listenner at IP: %s, Port: %d, is starting\n", s.IP, s.Port)

	go func() {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("Resolve tcp addr error:", err)
			return
		}

		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("Listen ", s.IPVersion, " error ", err)
			return
		}
		fmt.Println("Start server ", s.Name, "success, now listening...")

		var cid uint32 = 0
		// Server loops and waits for client connection
		for {
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept error:", err)
				continue
			}

			dealConn := NewConnection(conn, cid, s.Router)
			cid++

			go dealConn.Start()
		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("[STOP] Server ", s.Name)
}

func (s *Server) Serve() {
	s.Start()

	select {}
}

func (s *Server) AddRouter(router iface.IRouter) {
	s.Router = router
	fmt.Println("Add Router success")
}
