package main

import (
	"context"
	dockerController "docker_whisperer/docker_controller"
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
	dc, _ := dockerController.NewController()
	imagesList, _ := dc.ListImages()
	return &pb.GetDockerImagesListResponse{
		DockerImages: imagesList,
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
