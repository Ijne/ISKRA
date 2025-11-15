// package main

// import "fmt"

// func main() {
// 	fmt.Println("Here i am")

// 	// TODO: config loading

// 	// TODO: logger

// 	// TODO: storage

// 	// TODO: router

// 	// TODO: listen http
// }

package main

import (
	"iskra/bot"
	"iskra/shared/config"
)

func main() {
	cfg, err := config.New("./config/local.yaml")
	if err != nil {
		panic(err)
	}

	b := bot.New(cfg)
	done := b.ListenUpdates()
	<-done
}
