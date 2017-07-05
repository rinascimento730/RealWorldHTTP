package main

import (
    "io"
    "net/http"
    "os"
)

func main() {
    var url string = "http://wedding.gnavi.co.jp"
    var image_path string = "/Users/fujinaga/sandbox/go/app/images/"

    response, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    defer response.Body.Close()

    file, err := os.Create(image_path + "save.jpg")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    io.Copy(file, response.Body)
}