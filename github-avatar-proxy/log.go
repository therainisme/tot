package main

import (
	"fmt"
	"log"
)

type ServerResponse struct {
	Data interface{} `json:"data"`
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
}

type LogStore struct {
	Logs []string `json:"logs"`
}

func (s *LogStore) Write(p []byte) (n int, err error) {
	fmt.Println(string(p[:len(p)-1]))
	s.Logs = append(s.Logs, string(p[:len(p)-1]))
	return len(p), nil
}

var logStore = &LogStore{
	Logs: make([]string, 0),
}

func init() {
	log.SetOutput(logStore)
}