package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn_mq := connect("amqp://uwxaanrv:ny1KNNFVjWCNjU3zcc5hYu9CNiRUef7q@bee.rmq.cloudamqp.com/uwxaanrv")
	ch := createChannel(conn_mq)
	q := DeclareQ(ch)
	InitRedis()

	fmt.Println("Launching server...")

	ln, err := net.Listen("tcp", ":8081")
	defer ln.Close()
	// accept connection on port

	if err != nil {
		fmt.Println(err)
	}

	// run loop forever (or until ctrl-c)
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		go handleConnection(conn, ch, q)
	}
}
