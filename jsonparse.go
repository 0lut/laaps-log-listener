package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type singleLog struct {
	Timestamp string `json:"timestamp"`
	Log       string `json:"log"`
}

type logBlob struct {
	Sender  string      `json:"sender"`
	Logtype string      `json:"logtype"`
	Process string      `json:"process"`
	Logs    []singleLog `json:"logs"`
}

func handleJSON(conn net.Conn) {
	defer conn.Close()
	d := json.NewDecoder(conn)
	var m logBlob
	err := d.Decode(&m)
	if err != nil {
		panic(err)
	}

	fmt.Println(m)

}

func parseJSON(jsonStr string) logBlob {

	m := logBlob{}
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println(m.Sender)
	}
	return m

}
