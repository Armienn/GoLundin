package main

import (
	"io/ioutil"

	"encoding/json"
)

type Thread struct {
	ID          int
	Title       string
	MainMessage string
	Responses   []string
}

func loadThreads() []Thread {
	file, err := ioutil.ReadFile("threads.json")
	if err != nil {
		return nil
	}
	var threads []Thread
	err = json.Unmarshal(file, &threads)
	if err != nil {
		return nil
	}
	return threads
}