package main

import (
	"io"
	"log"
	"net/http"
)

type logWriter struct { 

}

func main() {
	resp, err := http.Get("http://google.com")
	if err!= nil {
        log.Fatal("Error:", err)
    }

	lw := logWriter{}

	io.Copy(&lw, resp.Body)
}

func (l *logWriter) Write(bs []byte) (n int, err error) {
    log.Println(string(bs))
	log.Println("just wrote this many bytes:", len(bs))
    return len(bs), nil
}