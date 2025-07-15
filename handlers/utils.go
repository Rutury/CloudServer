package handlers

import (
	"net/http"
	"net/url"
)

func setFlashCookie(w http.ResponseWriter, message string) {
	cookie := &http.Cookie{
		Name:  "flash",
		Value: message,
		Path:  "/",
	}
	http.SetCookie(w, cookie)
}

func getFlashCookie(w http.ResponseWriter, r *http.Request) string {
	cookie, err := r.Cookie("flash")
	if err != nil {
		return ""
	}

	msg, _ := url.QueryUnescape(cookie.Value)
	delCookie := &http.Cookie{
		Name:   "flash",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, delCookie)

	return msg
}
