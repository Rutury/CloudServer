package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		http.Error(w, "Cant read", http.StatusBadRequest)
		return
	}

	files := r.MultipartForm.File["files"]
	if len(files) == 0 {
		http.Error(w, "Choose files", http.StatusBadRequest)
		return
	}

	uploaded := 0
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			continue
		}
		defer file.Close()
		saveDir := filepath.Join("static", "cloud")
		os.MkdirAll(saveDir, os.ModePerm)
		dstPath := filepath.Join(saveDir, fileHeader.Filename)
		dst, err := os.Create(dstPath)
		if err != nil {
			continue
		}
		defer dst.Close()

		_, err = io.Copy(dst, file)
		if err == nil {
			uploaded++
		}
	}

	if uploaded == 0 {
		setFlashCookie(w, "Cant load any files")
	} else {
		setFlashCookie(w, fmt.Sprintf("Loaded: %d / %d", uploaded, len(files)))
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
