package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

const (
	resultFilePath = "./result/"
	formFilePath   = "./index.html"
)

func GetForm(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.Error(res, "Адрес не найден", http.StatusNotFound)
		return
	}

	if req.Method != http.MethodGet {
		res.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	http.ServeFile(res, req, formFilePath)
	res.WriteHeader(http.StatusOK)
}

func UploadFile(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		res.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	err := req.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	file, header, err := req.FormFile("myFile")
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := service.Convert(string(data))
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	ext := filepath.Ext(header.Filename)
	filename := fmt.Sprintf("%sresult_%d%s", resultFilePath, time.Now().UTC().UnixNano(), ext)
	os.Mkdir(resultFilePath, 0755)
	output, err := os.Create(filename)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
	defer output.Close()

	_, err = output.WriteString(result)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}

	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(res, result)
}
