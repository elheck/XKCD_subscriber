package main

import (
	"log"
	"os"
	"fmt"
	"sync"
	"time"
)

var rootPath string

func main() {
	rootPath := os.Args[1]
	fmt.Println("RootPath is", rootPath)
	var wg sync.WaitGroup
	dailyTicker := time.NewTicker(12 * time.Second)
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
			log.Println("No new comic")
		}
	}
}
