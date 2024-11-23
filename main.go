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
	//router.Use(cors.Default())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://cinema-production-7150.up.railway.app"},     // Ganti dengan URL frontend yang sesuai
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},                      // Metode HTTP yang diizinkan
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Accept"}, // Header yang diizinkan
		AllowCredentials: true,                                                          // Izinkan pengiriman cookies atau header Authorization
	}))

	routers.CustInitiator(router)
	routers.Cinema_hallInitiator(router)
	routers.MovieInitiator(router)
	routers.ShowtimeInitiator(router)
	routers.BookingInitiator(router)

	// router.Run(":8080")
	router.Run(":" + os.Getenv("PORT"))
}
