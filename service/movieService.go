package service

import (
	//"bytes"

	"context"
	"encoding/json"

	//	"errors"
	"fmt"
	"io"

	//"log"
	"projecto/database"
	"projecto/service/movie"

	//	"google.golang.org/grpc/status"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	//	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Server struct{}

type MovieItem struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Image       string             `bson:"image"`
	Director    string             `bson:"director"`
	Duration    string             `bson:"duration"`
}

var collection = database.ConnectDB()

func (s *Server) Create(ctx context.Context, req *movie.CreateReq) (*movie.CreateRes, error) {
	moviee := req.GetMovie()
	data := MovieItem{

		Name:        moviee.Name,
		Description: moviee.Description,
		Image:       moviee.Image,
		Director:    moviee.Director,
		Duration:    moviee.Duration,
	}
	result, err := collection.InsertOne(context.TODO(), data)

	if err != nil {
		fmt.Printf("%v ", err)
	}
	fmt.Println("this is create function from server")
	json.NewEncoder(io.MultiWriter()).Encode(result)
	fmt.Println("added to mongodb")
	return &movie.CreateRes{Movie: moviee}, nil
}

func (s *Server) Read(ctx context.Context, req *movie.ReadReq) (*movie.ReadRes, error) {
	id := req.GetId()
	fmt.Printf("this is id: %v \n", id)

	result := collection.FindOne(context.TODO(), bson.M{"id": id})
	data := MovieItem{}
	err := result.Decode(&data)
	if err != nil {
		fmt.Printf("error while decoding...%v", err)
	}
	response := &movie.ReadRes{Movie: &movie.Movie{
		Id:          id,
		Name:        data.Name,
		Description: data.Description,
		Image:       data.Image,
		Director:    data.Director,
		Duration:    data.Duration,
	}}
	fmt.Println("this is a read function from server")
	return response, nil
}

func (s *Server) Delete(ctx context.Context, req *movie.DeleteReq) (*movie.DeleteRes, error) {
	id := req.GetId()
	fmt.Printf("ID at Delete function %v", id)
	_, err := collection.DeleteOne(context.TODO(), bson.M{"id": id})
	if err != nil {
		fmt.Println("error while deleting the movie")
	}
	return &movie.DeleteRes{Success: true}, nil

}

func (s *Server) ReadAll(ctx context.Context, req *movie.ReadAllReq) (*movie.ReadAllRes, error) {
	data := MovieItem{}

	cursor, err := collection.Find(context.TODO(), bson.M{})
	fmt.Printf("cursor value\n %v \n\n", cursor.Err())

	if err != nil {
		fmt.Println("error while searching list of movies")
	}
	list := []*movie.Movie{}
	for cursor.Next(context.Background()) {
		err := cursor.Decode(&data)

		if err != nil {
			fmt.Printf("error occures while decoding the cursor data: %v", err)
		}
		list = append(list, &movie.Movie{
			Id:          []byte(data.Id.String()),
			Name:        data.Name,
			Description: data.Description,
			Image:       data.Image,
			Director:    data.Director,
			Duration:    data.Duration,
		})
	}

	return &movie.ReadAllRes{
		Movie: list,
	}, nil

}
