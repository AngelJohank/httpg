package main

import (
	"fmt"
	"io"
	"net/http"
)

func FetchAndWriteImage(url string, ch chan string) {
	res, err := http.Get(url)

	if err != nil {
		ch <- err.Error()
		return
	}

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

	// Write image
	imageName := GenerateImageName(res, url)
	WriteImage(imageName, imageData)

	ch <- fmt.Sprintf("Finished: %q", url)
}
