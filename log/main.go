package log

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Printf("Request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)

		err := readBody(r)
		if err != nil {
			log.Printf("Error reading request body: %s", err)
		}

		next.ServeHTTP(w, r)
	})
}

func readBody(r *http.Request) error {
	if r.Method == "POST" || r.Method == "PUT" {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			return err
		}

		// After reading the body, you need to replace it for further handlers
		r.Body = io.NopCloser(strings.NewReader(string(body)))

		log.Printf("Body: %s\n", string(body))

		return nil
	}

	return nil
}
