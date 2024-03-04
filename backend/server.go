package backend

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"telegram-bot-spotify/backend/database"
	"telegram-bot-spotify/backend/profile"
)

type server struct {
	profile.UnimplementedProfileRetrievalServiceServer
}

var port = flag.Int("port", 50051, "Server port")

func (s *server) GetProfileById(ctx context.Context, in *profile.ProfileByIdRequest) (*profile.Profile, error) {
	profile := &profile.Profile{
		Id:       in.Id,
		Username: in.Username,
	}
	return profile, nil
}

func (s *server) AddProfile(ctx context.Context, in *profile.Profile) (*profile.ErrorResponse, error) {
	err := database.AddProfileToDB(in.Id, in.Username)
	if err != nil {
		return nil, err
	}
	return &profile.ErrorResponse{Message: "Тебя занесли в базу данных."}, nil
}

func Init() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		fmt.Println("Failed to listen:", err)
		return
	}

	s := grpc.NewServer()
	srv := &server{}

	profile.RegisterProfileRetrievalServiceServer(s, srv)
	fmt.Println("Server started at port", *port)

	if err := s.Serve(lis); err != nil {
		fmt.Println("Failed to serve:", err)
		return
	}
}
