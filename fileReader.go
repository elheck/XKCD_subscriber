package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

const jsonPath = "info.json"

//UnmarshalJSON reads JSON file into interface
func UnmarshalJSON() interface{} {
	file, _ := ioutil.ReadFile(jsonPath)
	var data interface{}
	err := json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal("Cannot unmarshal the json ", err)
	}
	return data
}
