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
	singleJSON := json.Logs[0]
	apiKey, err := GetApiKeyOwner(singleJSON.Sender)
	fmt.Print(json)
	if err != nil {
		return logBlob{}, errors.New("Cannot find such user!")
	}
	for _, v := range json.Logs {
		v.Sender = apiKey

	}

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
		fmt.Print(data)
		jsonData, err = manipulateJSON(jsonData)
		if err != nil {
			fmt.Print("bulamadim")
			return
		}
		for _, v := range jsonData.Logs {

			data2, _ := json.Marshal(v)
			go sendData(ch, q.Name, data2, "application/json")
		}

		toSend := "Succesfully recieved!"

		conn.Write([]byte(toSend))
		return
	}
}
