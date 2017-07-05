package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("simpleget/main.go")
	if err != nil {
		panic(err)
	}
	resp, err := http.Post("http://localhost:18888", "text/plan", file)
	if err != nil {
		panic(err)
	}
	log.Println("Status: ", resp.Status)

	reader := strings.NewReader("テキスト")
	resp2, err := http.Post("http://localhost:18888", "text/plan", reader)
	if err != nil {
		panic(err)
	}
	log.Println("Status: ", resp2.Status)
}