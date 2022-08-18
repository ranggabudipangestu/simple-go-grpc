package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ranggabudipangestu/go-grpc-client/models"
	"google.golang.org/grpc"

	userHandler "github.com/ranggabudipangestu/go-grpc-client/handlers/user"
)

func main() {
	port := "8827"
	targetPort := "8957"
	conn, err := grpc.Dial(":"+targetPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect into grpc server \n")
	}

	user := models.NewUsersClient(conn)
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "server is running properly")
	})
	userHandler.CreateUserHandler(router, user)

	err = router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
