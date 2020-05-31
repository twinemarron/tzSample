package main

import (
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, "from tzSample page..")
}

func main() {
	server := &http.Server{
		Addr: "127.0.0.1:8002",
	}
	http.HandleFunc("/", index)
	server.ListenAndServe()
}
