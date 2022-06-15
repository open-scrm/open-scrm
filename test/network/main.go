package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		return
	}
	for {
		cc, err := ln.Accept()
		if err != nil {
			return
		}
		fmt.Println(cc.LocalAddr().String())
	}
}
