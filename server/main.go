package main

import (
	"os"
	"server/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func main() {
	port := os.Getnev("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())

	/*
	Endpoints
	1. POST instruments (C)
	2. GET instruments (R)
	3. UPDATE instruments (U)
	4. DELETE instruments (D)

	*/

	// create financial instrument endpoints
	router.POST("/instrument/create", routes.AddOrder)
	router.GET("/instrument", routes.GetOrders)
	router.GET("/instrument/:id/", routes.GetOrderById)
	router.PUT("/instrument/update/:id", routes.UpdateOrder)
	router.DELETE("/instrument/delete/:id", routes.DeleteOrder)

	// run server
	router.Run(":" + port)

}