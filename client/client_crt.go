package main

import (
    "crypto/tls"
    "log"
    "net/http"
    "net/http/httputil"
)

func main() {
    cert, err := tls.LoadX509KeyPair("../server/ssl/client.crt", "../server/ssl/client.key")
    if err != nil {
        panic(err)
    }

    // クライアントを作成
    client := &http.Client{
        Transport: &http.Transport{
            TLSClientConfig: &tls.Config{
                Certificates: []tls.Certificate{cert},
            },
        },
    }

    // 通信を行う
    resp, err := client.Get("https://localhost:18444")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    dump, err := httputil.DumpResponse(resp, true)
    if err != nil {
        panic(err)
    }

    log.Println(string(dump))
}