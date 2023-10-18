package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hello Argo</h1>\n <h2>by Ine</h2>"))
	})
	http.ListenAndServe(":8080", nil)
}
