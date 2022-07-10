package main

import (
	"context"
	pb "docker_whisperer/dockersinventory/pb"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedInventoryServer
}

func (s *server) GetDockerImagesList(ctx context.Context, in *pb.GetDockerImagesListRequest) (*pb.GetDockerImagesListResponse, error) {
	log.Printf("Received request: %v", in.ProtoReflect().Descriptor().FullName())
	return &pb.GetDockerImagesListResponse{
		DockerImages: GetDockerImagesSample(),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterInventoryServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func GetDockerImagesSample() []*pb.DockerImage {
	dockerImages := []*pb.DockerImage{
		{
			Image: "alpine",
			Cmd:   "echo",
		},
	}
	return dockerImages
}
