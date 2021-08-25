package services

import (
	"context"

	pb "github.com/glebnaz/go-platform-hello-world/pkg/pb/api/v1"
)

type Service struct {
}

func (s Service) GetPet(ctx context.Context, request *pb.GetPetRequest) (*pb.GetPetResponse, error) {
	panic("implement me")
}

func (s Service) PutPet(ctx context.Context, request *pb.PutPetRequest) (*pb.PutPetResponse, error) {
	panic("implement me")
}

func (s Service) DeletePet(ctx context.Context, request *pb.DeletePetRequest) (*pb.DeletePetResponse, error) {
	panic("implement me")
}

func NewService() pb.PetStoreServer {
	return Service{}
}
