package main

import (
	"fmt"
	"log"
	"net"
	"projecto/database"
	"projecto/service"
	"projecto/service/movie"

	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":5050")

	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	//	srv := service.Server{}
	movie.RegisterMovieServiceServer(server, &service.Server{})
	database.ConnectDB()
	fmt.Println("this is running")
	err = server.Serve(listen)
	fmt.Println("this is also running")
	if err != nil {
		fmt.Println("error while serving grpc server")
	}
	fmt.Println("grpc server started at port 5050")

}
