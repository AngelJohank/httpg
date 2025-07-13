package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	urls := os.Args[1:]
	result := make(chan string)

	for _, url := range urls {
		go downloadImg(url, result)
	}

	for range urls {
		fmt.Println(<-result)
	}
}

func downloadImg(url string, result chan string) {
	res, err := http.Get(url)
	if err != nil {
		result <- fmt.Sprintf("failed to get image from %v\n", url)
		return
	}

	resType := res.Header.Get("content-type")
	if !strings.HasPrefix(resType, "image/") {
		result <- "the given url does not contain any image"
		return
	}

	resBody, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		result <- fmt.Sprintf("failed to read image bytes from %v\n", url)
		return
	}

	urlSegments := strings.Split(url, "/")
	filename := urlSegments[len(urlSegments)-1]

	wd, err := os.Getwd()
	if err != nil {
		result <- "could not get current working directory"
	}

	err = os.WriteFile(filepath.Join(wd, filename), resBody, 0644)
	if err != nil {
		result <- err.Error()
	}

	result <- fmt.Sprintf("finished: %q", url)
}
