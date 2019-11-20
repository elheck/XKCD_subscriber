package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

//Info is the data type for the json data
type Info struct {
	Username    string   `json:"username"`
	Password    string   `json:"pw"`
	SMTPServer  string   `json:"smtp_server"`
	SMTPPort    string   `json:"smtp_port"`
	Addresslist []string `json:"addresslist"`
}

const jsonPath = "info.json"

//UnmarshalJSON reads JSON file into interface
func UnmarshalJSON() Info {
	file, _ := ioutil.ReadFile(jsonPath)
	var data Info
	err := json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal("Cannot unmarshal the json ", err)
	}
	return data
}
