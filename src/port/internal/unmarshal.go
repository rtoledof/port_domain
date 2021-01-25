package internal

import (
	"portdb.io/src/models"
	"portdb.io/src/proto/grpc"
)

func Unmarshal(pb *grpc.Port) *models.Port {
	return &models.Port{
		Code:        pb.Code,
		Name:        pb.Name,
		City:        pb.City,
		Country:     pb.Country,
		Alias:       pb.Alias,
		Regions:     pb.Regions,
		Coordinates: pb.Coordinates,
		Province:    pb.Province,
		Timezone:    pb.Timezone,
		Unlocs:      pb.Unlocs,
	}
}
