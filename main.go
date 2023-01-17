package main

import "net/http"

func RootHandler(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "http://www.google.com", http.StatusMovedPermanently)
}

func main() {

	http.HandleFunc("/", RootHandler)
	http.ListenAndServe(":8080", nil)

}
