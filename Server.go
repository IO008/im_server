package main

import (
	"fmt"
	mnet "im_server/net"
)

func main() {
	fmt.Println("Hello World!")
	mnet.NewServer("Hello").Start()
}
