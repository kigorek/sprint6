package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"path/filepath"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {

	indexfile := filepath.Join(".", "index.html")

	data, err := os.ReadFile(indexfile)

	if err != nil {
		log.Printf("Failed to read index.html: %v", err)
		http.Error(w, "Failed to load page", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(data)
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)

	if err != nil {
		log.Printf("Error parse form: %v", err)
		http.Error(w, "Error parse form", http.StatusInternalServerError)
		return
	}

	file, _, err := r.FormFile("myFile")

	if err != nil {
		log.Printf("error when uploading a file: %v", err)
		http.Error(w, "error when uploading a file", http.StatusInternalServerError)
		return
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Printf("error loading file: %v", err)
		http.Error(w, "error loading file", http.StatusInternalServerError)
		return
	}

	err = os.MkdirAll("result", 0755)

	if err != nil {
		log.Printf("error create dir: %v", err)
		http.Error(w, "error create dir", http.StatusInternalServerError)
		return
	}

	convertedText, _ := service.MorseOrText(string(data))
	timestamp := time.Now().UTC().Format("20060102-150405")

	fileName := filepath.Join("result", timestamp+".txt")
	err = os.WriteFile(fileName, []byte(convertedText), 0644)
	if err != nil {
		log.Printf("error saving to file: %v", err)
		http.Error(w, "error saving to file", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(convertedText))

}
