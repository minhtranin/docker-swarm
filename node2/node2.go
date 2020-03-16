package main

import (
    "fmt"
	"net/http"
	"os"
)

func main() {
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":8086", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	value,_ := os.Hostname();
    fmt.Fprintf(w, "iam golang node2, %s!", value)
}
