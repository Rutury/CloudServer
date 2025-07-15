package main

import (
	"CloudServer/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/upload", handlers.UploadHandler)
	http.HandleFunc("/delete", handlers.DeleteHandler)
	http.HandleFunc("/", handlers.SiteHandler)

	cloudFiles := http.FileServer(http.Dir("./static/cloud"))
	http.Handle("/cloud/", http.StripPrefix("/cloud/", cloudFiles))
	siteFiles := http.FileServer(http.Dir("./site"))
	http.Handle("/site/", http.StripPrefix("/site/", siteFiles))

	fmt.Println("Server running at http://localhost:5000")
	http.ListenAndServe(":5000", nil)
}
