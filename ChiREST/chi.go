package main

import (
	"flag"
	"github.com/TechMaster/LearnGo/ChiREST/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"encoding/json"
)

func main() {
	flag.Parse()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	//r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	r.Get("/articles", func(w http.ResponseWriter, r *http.Request) {
		b, _ := json.Marshal(models.GetArticles())
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)

	})

	http.ListenAndServe(":3333", r)
}






