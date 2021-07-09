package main

import (
	"context"
	"fmt"
	"log"
	"projecto/service/movie"
	"time"

	//	"go.mongodb.org/mongo-driver/bson/primitive"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":5050", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error while dialing port 5050: %v", err)
	}
	defer conn.Close()
	c := movie.NewMovieServiceClient(conn)
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	fmt.Println("client connected to server")

	req1 := movie.CreateReq{
		Movie: &movie.Movie{

			Name:        "The war of tomorrow",
			Description: "the retired military personal goes for the war which take place in the future",
			Image:       "this is image",
			Director:    "Chris Mckay",
			Duration:    "2h 20min",
		},
	}
	res1, err := c.Create(ctx, &req1)
	if err != nil {
		log.Fatalf("error while creating movie data: %v", err)
	}
	log.Printf("Create Result <%+v>\n\n", res1)

	id := res1.Id
	fmt.Printf("client id %v", id)

	req2 := movie.ReadReq{Id: id}
	res2, err := c.Read(ctx, &req2)
	if err != nil {
		log.Fatalf("error while reading movie data: %v\n\n", err)
	}
	log.Printf("Read Result <%+v>\n\n", res2)

	req3 := movie.ReadAllReq{}
	res3, err := c.ReadAll(ctx, &req3)

	if err != nil {
		log.Fatalf("ReadAll Failed:%v", err)
	}
	log.Printf("ReadAll Result\n <%+v> \n\n", res3)

	//req4 := movie.DeleteReq{Id: id}
	//res4, err := c.Delete(ctx, &req4)
	//if err != nil {
	//	log.Fatalf("Delete Failed: %v", err)
	//}
	//log.Printf("Delete Result <%+v>\n\n", res4)

}
