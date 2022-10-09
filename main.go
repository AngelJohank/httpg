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

	for range urls {
		fmt.Println(<-ch)
	}
}