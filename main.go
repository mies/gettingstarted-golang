package main

import (
	"os"
	"net/http"
	"encoding/json"
	"log"
)


func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal([]string{"San Francisco", "Amsterdam",
	"Berlin","Palo Alto", "Los Altos})
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(data)
}

func main() {
	http.HandleFunc("/", HomeHandler)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
