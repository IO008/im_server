package net

import (
	"fmt"
	"net"
)

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

		// Server loops and waits for client connection
		for {
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept error:", err)
				continue
			}

			go func() {
				for {
					buff := make([]byte, 512)
					count, err := conn.Read(buff)
					if err != nil {
						fmt.Println("Receive buffer error:", err)
						return
					}

					if _, err := conn.Write(buff[:count]); err != nil {
						fmt.Println("Write back buff error:", err)
						return
					}
				}
			}()
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
