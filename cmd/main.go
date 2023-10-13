package main

import (
	"log"

	"github.com/Nadeem1815/go-grpc-api-gateway/pkg/auth"
	"github.com/Nadeem1815/go-grpc-api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("failed at config", err)

	}
	r := gin.Default()
	auth.RegisterRoutes(r, &c)
	r.Run(c.Port)

}
