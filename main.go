package main

import (
	"log"
	"net"
	"os"
	"pea/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	cors "github.com/rs/cors/wrapper/gin"

	db "pea/configs"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func main() {
	r := gin.Default()
	//mongo connection
	db.Connect()

	// Getting host ip
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)

	r.Use(cors.Default())
	routes.SetupRoutes(r)
	err := r.Run((addrs[1].String()) + ":" + (os.Getenv("PORT")))
	if err != nil {
		log.Fatal(err)
	}
}
