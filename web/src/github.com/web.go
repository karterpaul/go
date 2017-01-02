package main

import (
	"io"
	"net/http"
        "log"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "web world!")
        log.Printf("reqval:%s",r.RemoteAddr)
        log.Printf("host:%s",r.Host)
        log.Printf("url:%s",r.URL)
        log.Printf("reqval:%s",r.Method)
        //log.Printf("reqval:%s",r.RemoteAddr)
        //log.Printf("reqval:%s",r.RemoteAddr)
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}

