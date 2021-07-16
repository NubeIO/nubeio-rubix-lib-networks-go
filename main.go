package main

import (
	"log"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/config"
	_ "github.com/NubeIO/nubeio-rubix-lib-rest-go/docs"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/helpers"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/pkg/database"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/pkg/router"
)

//func init() {
//	database.Setup()
//	helpers.DisableLogging(false)
//}
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
	database.Setup()
	helpers.DisableLogging(false)
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
