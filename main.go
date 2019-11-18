package main

const rootPath string = "./pics/"

func main() {
	title, path := SaveComicGetName(rootPath)
	SendMailToAddress("", path, title)
}
