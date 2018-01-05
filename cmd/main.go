package main

import (
	"runtime"
	"net/http"
	"encoding/json"
	"github.com/su-kun1899/meander"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// mendar.APIKey = "TODO"
	http.HandleFunc("/journeys", func(w http.ResponseWriter, r *http.Request){
		respond(w, r, meander.Journeys)
	})
	http.ListenAndServe(":8080", http.DefaultServeMux)
}

func respond(w http.ResponseWriter, r *http.Request, data []interface{}) error {
	return json.NewEncoder(w).Encode(data)
}
