package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/markbates/pkger"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		f, err := pkger.Open("/public/index.html")
		if err != nil {
			w.Write([]byte(err.Error()))
		}
		io.Copy(w, f)
		defer f.Close()
		return
	})

	dir := http.FileServer(pkger.Dir("/public"))
	r.PathPrefix("/public").Handler(http.StripPrefix("/public", dir))

	log.Fatal(http.ListenAndServe(":3000", r))
}
