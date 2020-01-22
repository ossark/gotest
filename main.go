package main

import (
	"fmt"
	"time"
	"net/http"
	"io/ioutil"
)

type item struct {
    value  string
    expiry time.Time
}

var data = make(map[string]*item)

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
    case "GET":
		if data[r.URL.Path] != nil { 
			fmt.Fprintf(w, data[r.URL.Path].value+"\n")
		} else { 
			w.WriteHeader(http.StatusNotFound) 
		} 
    case "POST":
    	body, err := ioutil.ReadAll(r.Body)
    	if err == nil { 
    	 	data[r.URL.Path] = &item { value: string(body), expiry : time.Now().Add(time.Second * 1800) }
    	}
    default:
        fmt.Fprintf(w, "GET or POST please.")
    }
}

func clean() {
	for {
		time.Sleep(60 * time.Second)
		for key, value := range data {
			if time.Now().After(value.expiry) {
				delete(data, key)
			}
		} 
	}
}

func main() {
	go clean()

	http.HandleFunc("/", handler)
	http.ListenAndServe(":1337", nil)
}