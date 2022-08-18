package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ranggabudipangestu/go-grpc-server/models"
	userHandler "github.com/ranggabudipangestu/go-grpc-server/modules/user/handler"
	userRepository "github.com/ranggabudipangestu/go-grpc-server/modules/user/repository"
	userService "github.com/ranggabudipangestu/go-grpc-server/modules/user/service"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	port := "8957"
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/grpc_server?parseTime=true"))

	if err != nil {
		log.Fatal(err)
	}

	db.Debug().AutoMigrate(
		models.UserDB{},
	)

	server := grpc.NewServer()

	userRepository := userRepository.NewUserRepository(db)
	userService := userService.NewUserService(userRepository)
	userHandler.NewUserHandler(server, userService)

	conn, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server Starting at Port: ", port)
	log.Fatal(server.Serve(conn))

}
