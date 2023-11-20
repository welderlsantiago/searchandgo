package router

import (
	"github.com/gin-gonic/gin" //Web Framework
)

func Initialize() {
	//Initializing Router
	router := gin.Default()

	//Initializing Routes
	initializeRoutes(router)

	//Server start
	router.Run(":8080") // listeningon 0.0.0.0:8080

}
