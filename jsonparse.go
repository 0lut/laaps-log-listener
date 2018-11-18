package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type singleLog struct {
	Timestamp string      `json:"timestamp"`
	Log       interface{} `json:"log"`
}

type logBlob struct {
	Sender string      `json: "sender"`
	Logs   []singleLog `json: "logs"`
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
func singleLogRet() singleLog {
	return singleLog{Timestamp: "123123", Log: []string{"/ GET 400", "/api POST 200"}}
}

func logRet() logBlob {
	return logBlob{Sender: "11", Logs: []singleLog{singleLogRet(), singleLogRet()}}
}

func q(jsonStr string) {

	b, _ := json.Marshal(logRet())
	m := logBlob{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		panic('a')
	} else {
		fmt.Println(m)
	}

}

// func main() {
// 	q("asd")

// }
