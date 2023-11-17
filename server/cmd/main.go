package main

import (
	"chat/server/db"
	"chat/server/internal/user"
	"chat/server/router"
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
	defer dbConn.Close()

	router.InitRouter(userHandler)
	router.Start("0.0.0.0:8080")
}
