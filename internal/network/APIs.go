package network

import "net/http"

func setupMux() *http.ServeMux {
	mux := http.NewServeMux()
	return mux
}
