package handlers

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	htmlContent, err := os.ReadFile("../index.html")
	if err != nil {
		http.Error(w, "fail to read index", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write(htmlContent)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 * 1024 * 1024)
	if err != nil {
		http.Error(w, "failed to parse form", http.StatusInternalServerError)
		return
	}

	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "fail to get file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "failed to read file content", http.StatusInternalServerError)
		return
	}

	if len(content) == 0 {
		http.Error(w, "file is empty", http.StatusInternalServerError)
		return
	}

	result, err := service.Translate(string(content))
	if err != nil {
		http.Error(w, "translation failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	filename := time.Now().UTC().Format("2006-01-02T15-04-05.999999999Z") + ".txt"
	outFile, err := os.Create(filename)
	if err != nil {
		http.Error(w, "failed to create output file", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	_, err = outFile.WriteString(result)
	if err != nil {
		http.Error(w, "failed to write result to file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
