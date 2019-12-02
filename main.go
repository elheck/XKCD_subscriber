package main

import "sync"

const rootPath string = "./pics/"

func main() {
	var wg sync.WaitGroup
	title, path := SaveComicGetName(rootPath)
	jsonInfo := UnmarshalJSON()
	server := jsonInfo.SMTPServer + ":" + jsonInfo.SMTPPort
	for _, address := range jsonInfo.Addresslist {
		wg.Add(1)
		go SendMailToAddress(address, path, title, jsonInfo.Username, jsonInfo.Password, server, jsonInfo.SMTPServer, &wg)
	}
	wg.Wait()
	//run once daily
	//with a Ticker
	//check for new comic first
}
