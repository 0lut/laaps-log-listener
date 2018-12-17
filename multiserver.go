package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/pkg/errors"

	"github.com/streadway/amqp"
)

func manipulateJSON(json logBlob) (logBlob, error) {
	apiKey, err := GetApiKeyOwner(json.Sender)
	fmt.Print(json)
	if err != nil {
		return logBlob{}, errors.New("Cannot find such user!")
	}
	json.Sender = apiKey

	return json, nil

}

func handleConnection(conn net.Conn, ch *amqp.Channel, q amqp.Queue) {
	fmt.Println("Serving to:", conn.RemoteAddr())
	defer conn.Close()
	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Println(err)
			log.Printf("Fail in ")
			return
		}
		jsonData := parseJSON(data)
		jsonData, err = manipulateJSON(jsonData)
		if err != nil {
			fmt.Print("bulamadim")
			return
		}
		dataBytes, _ := json.Marshal(jsonData)

		go sendData(ch, q.Name, dataBytes, "application/json")

		toSend := "Succesfully recieved!"

		conn.Write([]byte(toSend))
		return
	}
}
