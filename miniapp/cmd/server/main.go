package main

import (
	"fmt"
	"iskra/miniapp/internal/handlers"
	"iskra/miniapp/internal/middleware"
	"iskra/miniapp/internal/tools/timepad"
	"iskra/shared/config"
	"iskra/shared/storage/memgraph"
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
	g, err := memgraph.NewStorage(cfg)
	if err != nil {
		fmt.Println("Memgraph is not started")
		// panic(err)
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
	r.Use(middleware.CorsMiddleware("*"))
	r.Use(middleware.UserMiddleware())

	//r.Get("/rec-users", handlers.GetRecUsersHandler(s))
	r.Post("/like-user", handlers.LikeUserHandler(s, nil))
	r.Handle("/events", handlers.GetEventsHandler(s, t))
	r.Post("/flames", handlers.GetFlamesHandler(s))
	r.Post("/flame", handlers.CreateFlameHandler(s, t))
	r.Put("/flame", handlers.UpdateFlameHandler(s))
	r.Delete("/flame", handlers.DeleteFlameHandler(s))
	// r.Get("/user", handlers.GetProfileHandler(s)) // later change to "/profile"

	// r.Handle("/static/*", http.StripPrefix("/static/", fs))
	// r.Handle("/", http.HandlerFunc(handlers.StartScreenHandler(cfg)))

	// r.Handle("/static/*", http.StripPrefix("/static/", fs))
	// old routes
	// r.Handle("/start", http.HandlerFunc(handlers.StartScreenHandler(cfg)))
	// r.Handle("/", http.HandlerFunc(handlers.HomepageScreenHandler(cfg)))
	r.Handle("/profile", http.HandlerFunc(handlers.ProfileScreenHandler(cfg, s)))
	r.Handle("/createuser", http.HandlerFunc(handlers.CreateUserHandler(cfg, s, g)))
	r.Handle("/updateuser", http.HandlerFunc(handlers.UpdateUserHandler(cfg, s, g)))
	r.Handle("/recommendations", http.HandlerFunc(handlers.GetRecomendationsHandler(cfg, g)))
	r.Handle("/interaction", http.HandlerFunc(handlers.InteractionHandler(cfg, s, g)))

	// server
	log.Println("Start serving...")
	if err := http.ListenAndServe(cfg.Server.Address, r); err != nil {
		log.Println(err)
	}
}
