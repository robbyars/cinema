package main

import (
	"cinema/configs"
	"cinema/databases/connection"
	"cinema/databases/migration"
	"cinema/routers"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	configs.Initiator()

	connection.Initiator()
	defer connection.SqlDBConnections.Close()

	migration.Initiator(connection.SqlDBConnections)

	InitiateRouter()
}

func InitiateRouter() {
	router := gin.Default()
	router.Use(cors.Default())

	routers.CustInitiator(router)
	routers.Cinema_hallInitiator(router)
	routers.MovieInitiator(router)
	routers.ShowtimeInitiator(router)
	routers.BookingInitiator(router)

	// router.Run(":8080")
	router.Run(":" + os.Getenv("PORT"))
}
