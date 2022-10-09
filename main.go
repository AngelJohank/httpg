package main

import (
	"fmt"
	"os"
)

func main() {
	urls := ParseURLS(os.Args[1:])
	ch := make(chan string)

	for _, url := range urls {
		go FetchAndWriteImage(url, ch)
	}

	// Get all status messages
	for range urls {
		fmt.Println(<-ch)
	}
}