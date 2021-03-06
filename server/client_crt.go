package main

import (
    "crypto/tls"
    "fmt"
    "log"
    "net/http"
    "net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
    dump, err := httputil.DumpRequest(r, true)
    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        return
    }
    fmt.Println(string(dump))
    fmt.Fprintf(w, "<html><body>hello</body></html>\n")
}

func main() {
    server := &http.Server{
        TLSConfig: &tls.Config{
            ClientAuth: tls.RequireAndVerifyClientCert,
            MinVersion: tls.VersionTLS12,
        },
        Addr: ":18444",
    }
    http.HandleFunc("/", handler)
    log.Println("start http listening :18444")
    err := server.ListenAndServeTLS("ssl/server.crt", "ssl/server.key")
    log.Println(err)
}