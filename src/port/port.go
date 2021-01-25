package port

import (
	context "context"
	"fmt"
	"log"

	"github.com/google/uuid"

	"portdb.io/src/port/internal"

	"portdb.io/src/models"
	"portdb.io/src/proto/grpc"
)

var _ grpc.PortDomainServiceServer = &service{}

type service struct {
	repo models.PortDomainRepository
	grpc.UnimplementedPortDomainServiceServer
}

func (s *service) Store(ctx context.Context, request *grpc.CreateRequest) (*grpc.Port, error) {
	req := Unmarshal(request)
	if err := req.Valid(); err != nil {
		return nil, err
	}
	present := true
	port, _ := s.repo.Fetch(req.Code)
	if port.Code == "" {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, fmt.Errorf("unable to generate a new uuid")
		}
		port = models.Port{
			ID:          id.String(),
			Code:        req.Code,
			Name:        req.Name,
			City:        req.City,
			Country:     req.Country,
			Alias:       req.Alias,
			Regions:     req.Regions,
			Coordinates: req.Coordinates,
			Province:    req.Province,
			Timezone:    req.Timezone,
			Unlocs:      req.Unlocs,
		}
		present = false
	}
	var err error
	if !present {
		log.Printf("creating port %s", port.Name)
		err = s.repo.Store(&port)
	} else {
		log.Printf("updating port %s", port.Name)
		err = s.repo.Update(&port)
	}
	if err != nil {
		return nil, err
	}
	return internal.Marshal(port), nil
}

func (s *service) Fetch(ctx context.Context, request *grpc.FetchRequest) (*grpc.Port, error) {
	p, err := s.repo.Fetch(request.Code)
	if err != nil {
		return nil, err
	}
	return internal.Marshal(p), nil
}

func New(db models.PortDomainRepository) *service {
	return &service{repo: db}
}
