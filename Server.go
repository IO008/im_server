package main

import (
	"fmt"
	"im_server/iface"
	mnet "im_server/net"
)

type PingRouter struct {
	mnet.BaseRouter
}

func (pr *PingRouter) PreHandle(request iface.IRequest) {
	fmt.Println("Call Ping Router PreHandle")

	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping...\n"))
	if err != nil {
		fmt.Println("before ping error")
	}
}

func (pr *PingRouter) Handle(request iface.IRequest) {
	fmt.Println("Call Ping Router Handle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping\n"))
	if err != nil {
		fmt.Println("ping error")
	}
}

func (pr *PingRouter) PostHandle(request iface.IRequest) {
	fmt.Println("Call Ping Router PostHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping...\n"))
	if err != nil {
		fmt.Println("after ping error")
	}
}

func main() {
	server := mnet.NewServer("simple im ")
	server.AddRouter(&PingRouter{})
	server.Serve()
}
