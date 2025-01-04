package net

import (
	"fmt"
	"im_server/iface"
	"net"
)

type Connection struct {
	Conn    *net.TCPConn
	ConnID  uint32
	isClose bool

	Router iface.IRouter

	ExitBuffChan chan bool
}

func NewConnection(conn *net.TCPConn, connID uint32, router iface.IRouter) *Connection {
	return &Connection{
		Conn:         conn,
		ConnID:       connID,
		Router:       router,
		isClose:      false,
		ExitBuffChan: make(chan bool, 1),
	}
}

func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running")
	defer fmt.Println(c.RemoteAddr().String(), "conn reader exit!")
	defer c.Stop()

	for {
		buff := make([]byte, 512)
		_, err := c.Conn.Read(buff)
		if err != nil {
			fmt.Println("Receive buffer error:", err)
			c.ExitBuffChan <- true
			return
		}

		req := Request{
			conn: c,
			data: buff,
		}
		go func(request iface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
	}
}

func (c *Connection) Start() {
	go c.StartReader()

	for range c.ExitBuffChan {
		return
	}

}

func (c *Connection) Stop() {
	if c.isClose {
		return
	}
	c.isClose = true

	c.Conn.Close()

	c.ExitBuffChan <- true

	close(c.ExitBuffChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}
