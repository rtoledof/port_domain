package port

import (
	"portdb.io/src/proto/grpc"
	"portdb.io/src/validator"
)

var _ validator.Validator = &CreateRequest{}

type CreateRequest struct {
	Name        string
	City        string
	Country     string
	Alias       []string
	Regions     []string
	Coordinates []float32
	Province    string
	Timezone    string
	Unlocs      []string
	Code        string
}

func (r *CreateRequest) Valid() error {
	return nil
}

func Unmarshal(pb *grpc.CreateRequest) CreateRequest {
	return CreateRequest{
		Name:        pb.Name,
		City:        pb.City,
		Country:     pb.Country,
		Alias:       pb.Alias,
		Regions:     pb.Regions,
		Coordinates: pb.Coordinates,
		Province:    pb.Province,
		Timezone:    pb.Timezone,
		Unlocs:      pb.Unlocs,
		Code:        pb.Code,
	}
}
