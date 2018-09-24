package main

import (
	"flag"
	"github.com/TechMaster/LearnGo/ChiREST/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"encoding/json"
)

var routes = flag.Bool("routes", false, "Generate router documentation")

// Article fixture data
var articles = []*models.Article{
	{	ID:     "1",
		UserID: 1,
		Title:  "Multi-layered 5th generation workforce",
		Slug:   "radical" },
	{	ID:     "2",
		UserID: 6,
		Title:  "Future-proofed 6th generation secured line",
		Slug:   "systematic"},
	{
		ID:     "3",
		UserID: 4,
		Title:  "Automated bandwidth-monitored matrix",
		Slug:   "Implemented"},
	{
		ID:     "4",
		UserID: 3,
		Title:  "Balanced human-resource help-desk",
		Slug:   "frame"},
	{
		ID:     "5",
		UserID: 7,
		Title:  "Secured leading edge circuit",
		Slug:   "systemic"},
	{
		ID:     "6",
		UserID: 2,
		Title:  "Public-key dynamic matrix",
		Slug:   "tertiary"},
	{
		ID:     "7",
		UserID: 8,
		Title:  "Triple-buffered attitude-oriented access",
		Slug:   "high-level"},
	{
		ID:     "8",
		UserID: 2,
		Title:  "Up-sized content-based installation",
		Slug:   "Polarised"},
	{
		ID:     "9",
		UserID: 8,
		Title:  "Monitored even-keeled infrastructure",
		Slug:   "asynchronous"},
	{
		ID:     "10",
		UserID: 9,
		Title:  "Open-source object-oriented application",
		Slug:   "contextually-based"},
}


// User fixture data
var users = []*models.User{
	{ID: 1, Name: "Peter"},
	{ID: 2, Name: "Julia"},
	{ID: 3, Name: "Rock"},
	{ID: 4, Name: "Queen"},
	{ID: 5, Name: "Sydney"},
	{ID: 6, Name: "Tom Fort"},
	{ID: 7, Name: "Peterson Hans"},
	{ID: 8, Name: "Julia Racker"},
	{ID: 9, Name: "Joe Candis"},
}
// Sets the content type of response. Also adds the Server header.
func setContentType(w http.ResponseWriter, contentType string) {
	w.Header().Set("Server", "Chi")
	w.Header().Set("Content-Type", contentType)
}

func main() {
	flag.Parse()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	//r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("HelloWorld"))
	})

	r.Get("/articles", func(w http.ResponseWriter, r *http.Request) {
		b, _ := json.Marshal(articles)
		setContentType(w, "application/json")
		w.Write(b)
	})


	http.ListenAndServe(":3333", r)
}






