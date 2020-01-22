package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

var data = make(map[string]string)

func handler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
    case "GET":
		if data[r.URL.Path] != "" { 
			fmt.Fprintf(w, data[r.URL.Path]) 
		} else { 
			w.WriteHeader(http.StatusNotFound) 
	} 
    case "POST":
    	body, err := ioutil.ReadAll(r.Body)
    	if err == nil { 
    		data[r.URL.Path] = string(body) 
    	}
    default:
        fmt.Fprintf(w, "GET or POST please.")
    }
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":1337", nil)
}