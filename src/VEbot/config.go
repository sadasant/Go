package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Nick     string
	Password string
	Server   string
	Channel  string
}

func readConfig(filename string) {
	defer println("Reading Config at: " + filename)
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		println(err)
		return
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		println(err)
		return
	}
}
