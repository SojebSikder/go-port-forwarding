package main

import (
	"flag"
	"fmt"
	"io"
	"net"
)

var (
	target string
	port   string
)

func init() {
	flag.StringVar(&target, "target", "", "the target (<host>:<port>)")
	flag.StringVar(&port, "port", "3000", "the tunnelthing port")
}

func main() {
	flag.Parse()

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	fmt.Println("new client")

	proxy, err := net.Dial("tcp", target)
	if err != nil {
		panic(err)
	}

	fmt.Println("proxy connected")
	go copyIO(conn, proxy)
	go copyIO(proxy, conn)
}

func copyIO(src, dest net.Conn) {
	defer src.Close()
	defer dest.Close()
	io.Copy(src, dest)
}
