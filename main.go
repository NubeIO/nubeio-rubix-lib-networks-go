package main

import (
	"log"
	"rubix-lib-rest-go/config"
	"rubix-lib-rest-go/helpers"
	"rubix-lib-rest-go/pkg/database"
	"rubix-lib-rest-go/pkg/router"

	_ "rubix-lib-rest-go/docs"
)

func init() {

	database.Setup()
	helpers.DisableLogging(false)
	//models.ConnectDatabase(rubixPath)

}

// @title GO Nube API
// @version 1.0
// @description nube api docs
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8081

func main() {
	commonConfig := config.CommonConfig()
	db := database.GetDB()
	r := router.Setup(db)

	port := commonConfig.Server.Port
	err := r.Run("localhost:" + port)
	log.Printf("Server is starting at 127.0.0.1:%s",port)
	if err != nil {
		log.Printf("Server error %s", port)
		return
	}
}