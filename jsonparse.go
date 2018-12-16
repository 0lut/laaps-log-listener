package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type singleLog struct {
	Timestamp string   `json:"timestamp"`
	Log       []string `json:"log"`
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

// func singleLogRet() singleLog {
// 	return singleLog{Timestamp: "1889931", Log: []string{"/ qq 400", "/courier ROUTE-31.ugur 200"}}
// }

// func logRet() logBlob {
// 	return logBlob{Sender: "7881", Logs: []singleLog{singleLogRet(), singleLogRet()}}
// }

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

// func main() {
// 	parseJSON(`{"sender":"myApiKey","logs":[{"log":["my other log ","/hey","hello world"],"timestamp":"1544983681585"},{"log":["my other log ","/hey","hello world"],"timestamp":"1544983741610"},{"log":["my other log ","/hey","hello world"],"timestamp":"1544983742353"},{"log":["my other log ","/hey","hello world"],"timestamp":"1544983743016"},{"log":["my other log ","/hey","hello world"],"timestamp":"1544983743638"},{"log":["myother log ","/hey","hello world"],"timestamp":"1544983744326"}]}`)
// }
