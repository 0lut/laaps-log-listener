package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	fmt.Println("Serving to:", conn.RemoteAddr())
	defer conn.Close()
	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println((data))

		// if data == "STOP" {
		// 	return
		// }
		toSend := "Pinging back... " + data
		// toSend, _ := json.Marshal(`{"ugur": "1"}`)

		conn.Write([]byte(toSend))
		return
	}
}

func main() {

	fmt.Println("Launching server...")

	// listen on all interfaces
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

		go handleJSON(conn)
	}
}
