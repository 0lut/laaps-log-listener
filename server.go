package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	// only needed below for sample processing
)

func main() {

	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")

	// accept connection on port

	conn, err := ln.Accept()

	if err != nil {
		fmt.Println(err)
	}

	// run loop forever (or until ctrl-c)
	for {
		// will listen for message to process ending in newline (\n)
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {

			os.Exit(1)

		} else {

			HandleMessage(conn, message)
		}

		// fmt.Print("Message Received:", string(message))
		// sample process for string received
		// newmessage := strings.ToUpper(message)
		// send new string back to client
		// conn.Write([]byte(newmessage + "girdi mi"))

	}
}

func HandleMessage(conn net.Conn, msg string) {
	defer conn.Close()
	fmt.Println((conn.RemoteAddr()))

	fmt.Println("Receieved message: ", msg)

	newmsg := strings.ToUpper(msg) + "...Sending Back"

	conn.Write([]byte(newmsg))
}

func handleConnection(conn net.Conn) {
	fmt.Println("New connection established!", conn)
}
