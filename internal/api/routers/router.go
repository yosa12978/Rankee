package routers

import "net/http"

func NewMainRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write([]byte("{\"message\": \"Hello World!\"}"))
	})
	return mux
}
