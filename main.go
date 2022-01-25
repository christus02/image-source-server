package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	TULIP_FILENAME   = "tulip.jpeg"
	IMAGES_DIRECTORY = "images/"
)

var (
	PORT string
)

func init() {
	found, ok := os.LookupEnv("SERVER_PORT")
	if ok {
		PORT = found
	} else {
		PORT = "8080"
	}
}

func tulip(w http.ResponseWriter, req *http.Request) {
	fileBytes, err := ioutil.ReadFile(IMAGES_DIRECTORY + TULIP_FILENAME)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
	return
}

func main() {
	fmt.Println("HTTP Server Started and Running on PORT:", PORT)
	http.HandleFunc("/tulip", tulip)

	http.ListenAndServe(":"+PORT, nil)

}
