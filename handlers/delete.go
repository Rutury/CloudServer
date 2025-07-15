package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	r.ParseForm()
	paths := r.Form["paths"]
	if len(paths) == 0 {
		setFlashCookie(w, "Choose files")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	saveDir := filepath.Join("static", "cloud")
	os.MkdirAll(saveDir, os.ModePerm)

	deleted := 0
	for _, relPath := range paths {
		fullPath := filepath.Join(saveDir, relPath)
		if !filepathHasPrefix(fullPath, saveDir) {
			continue
		}

		err := os.Remove(fullPath)
		if err == nil || os.IsNotExist(err) {
			deleted++
		}
	}

	if deleted == 0 {
		setFlashCookie(w, "Cant delete any files")
	} else {
		setFlashCookie(w, fmt.Sprintf("Deleted: %d / %d", deleted, len(paths)))
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func filepathHasPrefix(path, prefix string) bool {
	absPath, _ := filepath.Abs(path)
	absPrefix, _ := filepath.Abs(prefix)
	return len(absPath) >= len(absPrefix) && absPath[:len(absPrefix)] == absPrefix
}
