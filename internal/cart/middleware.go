package cart

import (
	"log"
	"net/http"
)

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		next.ServeHTTP(w, r)
	})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method + ": " + r.RequestURI)

		next.ServeHTTP(w, r)
	})
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func authMiddleware(next http.Handler) http.Handler {
	// This should go in a database, but for simplicty, we are going to hardcode
	// a couple of them
	allowedUsers := make(map[string]string)

	allowedUsers["admin"] = "admin"
	allowedUsers["user"] = "user"

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()

		if !ok || allowedUsers[user] != pass {
			ErrorUnauthorised(w, "")
			return
		}

		next.ServeHTTP(w, r)
	})
}
