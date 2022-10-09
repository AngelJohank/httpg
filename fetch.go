package main

import (
	"io"
	"net/http"
)

func FetchAndWriteImage(url string, ch chan string) {
	res, err := http.Get(url)

	if err != nil {
		ch <- err.Error()
		return
	}

	// Check if url is pointing to an image
	if !URLisImage(res) {
		ch <- "URL is not an image"
		return
	}

	// Read image data
	imageData, err := io.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		ch <- err.Error()
		return
	}

	// Write image to disk
	imageName := GenerateImageName(res, url)
	WriteImage(imageName, imageData)

	ch <- "FINISHED: " + url
}
