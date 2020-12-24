package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", handleMain)
	http.ListenAndServe(":8001", nil)
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	data := r.FormValue("data")
	if (len(data) == 0) {
		w.Write([]byte("i am working"))
	}
	w.Write([]byte(data))
}
