package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bhushan-aruto/ice_cream_shop_rest_api/config"
	"github.com/bhushan-aruto/ice_cream_shop_rest_api/repository"
	"github.com/bhushan-aruto/ice_cream_shop_rest_api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {

	serverMode := os.Getenv("SERVER_MODE")

	if serverMode == "dev" {

		if err := godotenv.Load(".env"); err != nil {
			log.Fatalf("Unable to load .env file: %v", err)
		}
		log.Println(".env file loaded successfully.")

	} else if serverMode != "prod" {

		log.Fatalln("Invalid SERVER_MODE. Please set SERVER_MODE to 'dev' or 'prod'.")
	}
}

func main() {

	config.ConnectMongo()

	repository.InitFlavourRepo()

	router := gin.Default()

	routes.SetRoutes(router)

	fmt.Println("server started running on port 8080")
	log.Fatalln(http.ListenAndServe(":8080", router))

}
