package main

import (
	"encoding/json"
	"fmt"
)

type singleLog struct {
	Timestamp string `json:"timestamp"`
	Log       string `json:"log"`
	Sender    string `json:"sender"`
	Logtype   string `json:"logtype"`
	Process   string `json:"process"`
}

type logBlob struct {
	Logs []singleLog `json:"logs"`
}

func parseJSON(jsonStr string) logBlob {

	m := logBlob{}
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println(m.Logs[0])
	}
	return m

}
