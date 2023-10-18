package main

import (
	"fmt"
	"github.com/MatheusVict/User-Register-GO/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error on .env file")
	}
	fmt.Println(os.Getenv("TEST"))

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatalln("Error on start application: ", err)
	}
}
