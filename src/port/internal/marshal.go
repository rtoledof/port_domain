package internal

import (
	"portdb.io/src/models"
	"portdb.io/src/proto/grpc"
)

func Marshal(p models.Port) *grpc.Port {
	return &grpc.Port{
		Code:        p.Code,
		Name:        p.Name,
		City:        p.City,
		Country:     p.Country,
		Alias:       p.Alias,
		Regions:     p.Regions,
		Coordinates: p.Coordinates,
		Province:    p.Province,
		Timezone:    p.Timezone,
		Unlocs:      p.Unlocs,
	}
}
