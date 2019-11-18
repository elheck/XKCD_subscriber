package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/nishanths/go-xkcd"
)

//SaveComicGetName returns Comic Name and path
//and save file as number.png
func SaveComicGetName(root string) (string, string) {
	comic := getLatestComic()
	numberString := strconv.Itoa(comic.Number)
	path := getPath(numberString, root)
	saveComicToFile(comic.ImageURL, path)
	return comic.SafeTitle, path
}

func getLatestComic() xkcd.Comic {
	client := xkcd.NewClient()
	comic, err := client.Latest()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Comic URL:", comic.ImageURL)
	log.Println("Comic Name:", comic.SafeTitle)
	return comic
}

func getPath(name, root string) string {
	fileName := root + name + ".png"
	return fileName
}

func saveComicToFile(url, path string) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	file, er := os.Create(path)
	if er != nil {
		log.Fatal(er)
	}
	defer file.Close()
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Picture saved at: ", path)
}
