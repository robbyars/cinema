package main

import (
	"cinema/configs"
	"cinema/databases/connection"
	"cinema/databases/migration"
	"cinema/routers"

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

	routers.CustInitiator(router)
	routers.Cinema_hallInitiator(router)
	routers.MovieInitiator(router)
	routers.ShowtimeInitiator(router)
	routers.BookingInitiator(router)

	router.Run(":8080")
}
