package route

import (
	"fmt"
	"net/http"
	"url-shorter/server"

	"github.com/gorilla/mux"
)

func Route() *mux.Router {
	fmt.Println("In route folder main fn")

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is running on port 4001"))
	})
	r.HandleFunc("/api/allUrl", server.GetAllDocuments).Methods("GET")
	return r
}
