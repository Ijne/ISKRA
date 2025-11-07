package main

import (
	"fmt"
	"iskra/miniapp/internal/handlers"
	"iskra/shared/config"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func main() {
	// if err := godotenv.Load(); err != nil {
	// 	panic(err)
	// }
	// HOST := os.Getenv("HOST")
	// PORT := os.Getenv("PORT")

	// config
	cfg, err := config.New("./config/local.yaml")
	if err != nil {
		panic(err)
	}

	// static
	workDir, _ := os.Getwd()
	log.Printf("WD for loading static: %s\n", workDir)
	staticDir := filepath.Join(workDir, cfg.MiniApp.StaticPath)
	fs := http.FileServer(http.Dir(staticDir))

	// router
	r := chi.NewRouter()
	r.Handle("/static/*", http.StripPrefix("/static/", fs))
	r.Handle("/start", http.HandlerFunc(handlers.StartScreenHandler(cfg)))
	r.Handle("/", http.HandlerFunc(handlers.HomepageScreenHandler(cfg)))
	r.Handle("/profile", http.HandlerFunc(handlers.ProfileScreenHandler(cfg)))

	// server
	log.Println("Start serving...")
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.MiniApp.Host, cfg.MiniApp.Port), r); err != nil {
		log.Println(err)
	} else {
		log.Printf("Server started on %s:%s\n", cfg.MiniApp.Host, cfg.MiniApp.Port)
	}
}
