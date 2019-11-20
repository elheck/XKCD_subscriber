package main

const rootPath string = "./pics/"

func main() {
	title, path := SaveComicGetName(rootPath)
	jsonInfo := UnmarshalJSON()
	server := jsonInfo.SMTPServer + ":" + jsonInfo.SMTPPort
	for _, address := range jsonInfo.Addresslist {
		go SendMailToAddress(address, path, title, jsonInfo.Username, jsonInfo.Password, server, jsonInfo.SMTPServer)
	}

	//run once daily
	//with a Ticker
	//check for new comic first
}
