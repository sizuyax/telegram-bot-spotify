package backend

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	pb "telegram-bot-spotify/backend/profile"
)

func Client(id int64, username string) (*pb.ErrorResponse, error) {

	conn, err := grpc.Dial("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewProfileRetrievalServiceClient(conn)

	myProfile := &pb.Profile{
		Id:       id,
		Username: username,
	}

	resp, err := client.AddProfile(context.Background(), myProfile)
	if err != nil {
		log.Println("Failed to send gRPC request:", err)
		return nil, err
	}

	return resp, nil
}
