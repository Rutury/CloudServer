package handlers

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

func SiteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("site/site.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	saveDir := filepath.Join("static", "cloud")
	os.MkdirAll(saveDir, os.ModePerm)

	entries, err := os.ReadDir(saveDir)
	if err != nil {
		http.Error(w, "Read error", http.StatusInternalServerError)
		return
	}

	var metadata []FileData
	for _, entry := range entries {
		if !entry.IsDir() {
			metadata = append(metadata, FileData{entry.Name(), "img", "Some time ago"})
		}
	}
	msg := getFlashCookie(w, r)
	tmpl.Execute(w, SiteData{msg, metadata})
}

type FileData struct {
	Name       string
	Type       string
	UploadTime string
}

type SiteData struct {
	Msg   string
	Files []FileData
}
