package main

import (
	"github.com/jdwillmsen/go-chat-server/db"
	"github.com/jdwillmsen/go-chat-server/internal/router"
	"github.com/jdwillmsen/go-chat-server/internal/user"
	"github.com/jdwillmsen/go-chat-server/internal/ws"
	"log"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouter(userHandler, wsHandler)
	router.Start("0.0.0.0:8080")
}
