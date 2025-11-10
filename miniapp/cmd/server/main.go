package main

import (
	"fmt"
	"iskra/miniapp/internal/handlers"
	"iskra/miniapp/internal/tools/timepad"
	"iskra/shared/config"
	"iskra/shared/storage/postgres"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	// config
	cfg, err := config.New("./config/local.yaml")
	if err != nil {
		panic(err)
	}

	// storage
	s, err := postgres.NewStorage(cfg)
	if err != nil {
		panic(err)
	}

	// static
	// workDir, _ := os.Getwd()
	// log.Printf("WD for loading static: %s\n", workDir)
	// staticDir := filepath.Join(workDir, cfg.MiniApp.StaticPath)
	// fs := http.FileServer(http.Dir(staticDir))

	// timepad api
	t := timepad.New(cfg)

	// router
	r := chi.NewRouter()

	// new routes
	// r.Use(middleware.JWTAuthMiddleware(cfg))

	r.Get("/rec-users", handlers.GetRecUsersHandler(s))
	r.Post("/like-user", handlers.LikeUserHandler(s))
	r.Handle("/events", handlers.GetEventsHandler(s, t))
	r.Post("/flames", handlers.GetFlamesHandler(s))
	r.Post("/flame", handlers.CreateFlameHandler(s, t))
	r.Put("/flame", handlers.UpdateFlameHandler(s))
	r.Delete("/flame", handlers.DeleteFlameHandler(s))
	r.Get("/user", handlers.GetProfileHandler(s)) // later change to "/profile"

	// r.Handle("/static/*", http.StripPrefix("/static/", fs))
	// r.Handle("/", http.HandlerFunc(handlers.StartScreenHandler(cfg)))

	// r.Handle("/static/*", http.StripPrefix("/static/", fs))
	// old routes
	r.Handle("/start", http.HandlerFunc(handlers.StartScreenHandler(cfg)))
	r.Handle("/", http.HandlerFunc(handlers.HomepageScreenHandler(cfg)))
	r.Handle("/profile", http.HandlerFunc(handlers.ProfileScreenHandler(cfg)))
	r.Handle("/createuser", http.HandlerFunc(handlers.CreateUserHandler(cfg)))
	r.Handle("/updateuser", http.HandlerFunc(handlers.UpdateUserHandler(cfg)))

	// server
	log.Println("Start serving...")
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.MiniApp.Host, cfg.MiniApp.Port), r); err != nil {
		log.Println(err)
	} else {
		log.Printf("Server started on %s:%s\n", cfg.MiniApp.Host, cfg.MiniApp.Port)
	}
}
