package miniapp

import (
	"context"
	"iskra/bot"
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

type Miniapp struct {
	router  http.Handler
	bot     *bot.Bot
	cfg     *config.Config
	storage *postgres.Storage
	timepad *timepad.TimepadPoller
	ctx     context.Context
}

func New(cfg *config.Config, s *postgres.Storage, g *memgraph.Storage, bot *bot.Bot) *Miniapp {
	ctx, _ := context.WithCancel(context.Background())

	// timepad api
	t := timepad.New(cfg)

	// router
	r := chi.NewRouter()

	// new routes
	// r.Use(middleware.JWTAuthMiddleware(cfg))
	r.Use(middleware.CorsMiddleware("*"))
	r.Use(middleware.UserMiddleware())

	//r.Get("/rec-users", handlers.GetRecUsersHandler(s))
	r.Post("/like-user", handlers.LikeUserHandler(s, bot))
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
	r.Handle("/profile", http.HandlerFunc(handlers.ProfileScreenHandler(cfg, s)))
	r.Handle("/createuser", http.HandlerFunc(handlers.CreateUserHandler(cfg, s, g)))
	r.Handle("/updateuser", http.HandlerFunc(handlers.UpdateUserHandler(cfg, s, g)))
	r.Handle("/recommendations", http.HandlerFunc(handlers.GetRecomendationsHandler(cfg, g)))
	r.Handle("/interaction", http.HandlerFunc(handlers.InteractionHandler(cfg, s, g)))

	miniapp := Miniapp{
		router:  r,
		bot:     bot,
		cfg:     cfg,
		storage: s,
		timepad: t,
		ctx:     ctx,
	}

	return &miniapp
}

func (m *Miniapp) Listen() chan bool {
	done := make(chan bool)
	go func() {
		log.Println("Start serving...")
		if err := http.ListenAndServe(m.cfg.Server.Address, m.router); err != nil {
			log.Println(err)
		}
		done <- true
	}()
	return done
}
