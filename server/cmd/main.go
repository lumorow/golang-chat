package main

import (
	"chat/server/db"
	"chat/server/internal/user"
	"chat/server/internal/ws"
	"chat/server/router"
	"log"
)

func main() {
	dbConn, err := db.NewDatabase()
	defer dbConn.Close()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)

	router.InitRouter(userHandler, wsHandler)
	router.Start("0.0.0.0:8080")
}
