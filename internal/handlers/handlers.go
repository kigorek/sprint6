package handlers

import "net/http"

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	s := ""
	w.Write([]byte(s))
}

func UploadHandle(w http.ResponseWriter, r *http.Request) {
	s := ""
	w.Write([]byte(s))
}
