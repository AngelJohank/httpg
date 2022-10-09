package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func WriteImage(name string, data []byte) {
	err := os.WriteFile(name, data, 0644)

	if err != nil {
		fmt.Println(err)
	}
}

func URLisImage(res *http.Response) bool {
	contentType := res.Header.Get("content-type")
	return strings.HasPrefix(contentType, "image")
}

func GenerateImageName(res *http.Response, url string) string {
	contentType := res.Header.Get("content-type")

	imageName := base64.URLEncoding.EncodeToString([]byte(url))
	imageExtension := strings.Split(contentType, "/")[1]

	return fmt.Sprintf("%v.%v", imageName, imageExtension)
}