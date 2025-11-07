package main

import (
	"fmt"
	"iskra/internal/handlers"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	HOST := os.Getenv("HOST")
	PORT := os.Getenv("PORT")

	workDir, _ := os.Getwd()
	staticDir := filepath.Join(workDir, "internal/static")
	fs := http.FileServer(http.Dir(staticDir))

	r := chi.NewRouter()
	r.Handle("/static/*", http.StripPrefix("/static/", fs))
	r.Handle("/", http.HandlerFunc(handlers.StartScreenHandler))

	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", HOST, PORT), r); err != nil {
		log.Println(err)
	} else {
		log.Printf("Server started on %s:%s\n", HOST, PORT)
	}
}
