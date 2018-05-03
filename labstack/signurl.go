package main

import (
"fmt"
"time"

"cloud.google.com/go/storage"
)

const (
	projectID = "myProject-1234"
)

func main() {

	bucket := "mybucket"
	filename := "myFilename"
	method := "PUT"
	expires := time.Now().Add(time.Second * 60)

	url, err := storage.SignedURL(bucket, filename, &storage.SignedURLOptions{
		GoogleAccessID: "XXXXXX@developer.gserviceaccount.com",
		PrivateKey:     []byte("-----BEGIN PRIVATE KEY-----\nXXXXXXXX"),
		Method:         method,
		Expires:        expires,
	})
	if err != nil {
		fmt.Println("Error " + err.Error())
	}
	fmt.Println("URL = " + url)
}