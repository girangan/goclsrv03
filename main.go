package main

import (
	"fmt"
	"net/http"
    "time"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Hello from Cisco Shipped Pre-Existing Repo!</h1>\n")
    
    if !keepBusy{
        keepBusy = true
        fmt.Println("Keeping the process busy", time.Now())
        time.Sleep(1 * time.Second)
    }
}

var keepBusy bool

func main() {
    http.HandleFunc("/", defaultHandler)
    fmt.Println("Example app listening at http://localhost:8888")
    http.ListenAndServe(":8888", nil)
}

