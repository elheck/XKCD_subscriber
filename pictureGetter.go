package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/nishanths/go-xkcd"
)

//SaveComicGetName returns Comic Name and path
//and save file as number.png
func SaveComicGetName(root string) (string, string, bool) {
	comic := getLatestComic()
	numberString := strconv.Itoa(comic.Number)
	path := getPath(numberString, root)
	alreadyExists := false
	absPath, _ := filepath.Abs(path)
	_, err := os.Stat(absPath)
	if err == nil {
		alreadyExists = true
	} else if os.IsNotExist(err) {
		saveComicToFile(comic.ImageURL, path)
	}
	return comic.SafeTitle, path, alreadyExists
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
