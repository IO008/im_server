package main

import (
	mnet "im_server/net"
)

func main() {
	mnet.NewServer("simple im ").Serve()
}
