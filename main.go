package main

import (
	"github/jaseelaali/orchid/database"
	_ "github/jaseelaali/orchid/docs"
	"github/jaseelaali/orchid/routes"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	database.DatabaseConnection()
}
func main() {

	// @title ORCHID_FOOTWARES
	// @version 1.0
	// @description <add description here>

	// @contact.name API Support
	// @contact.url http://www.swagger.io/support
	// @contact.email support@swagger.io

	///@license.name MIT
	////@host localhost:9090
	//@host jaseela.tech
	// @license.url https://opensource.org/licenses/MIT

	//@host jaseela.tech
	////@host localhost:9090
	// @BasePath /
	// @query.collection.format multi

	Route := gin.Default()

	// swagger docs
	Route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	routes.User(Route)
	routes.Admin(Route)

	Route.Run(":9090")

}
