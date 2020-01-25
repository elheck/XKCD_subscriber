package main

import (
	"log"
	"sync"
	"time"
)

const rootPath string = "./pics/"

func main() {
	var wg sync.WaitGroup
	dailyTicker := time.NewTicker(24 * time.Hours)
	for _ = range dailyTicker.C {
		title, path, alreadyExists := SaveComicGetName(rootPath)
		if !alreadyExists{
			jsonInfo := UnmarshalJSON()
			server := jsonInfo.SMTPServer + ":" + jsonInfo.SMTPPort
			for _, address := range jsonInfo.Addresslist {
				wg.Add(1)
				go SendMailToAddress(address, path, title, jsonInfo.Username, jsonInfo.Password, server, jsonInfo.SMTPServer, &wg)
			}
			wg.Wait()
		} else {
			log.Println("No new comic today")
		}
	}
}
