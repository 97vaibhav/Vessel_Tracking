package handler

import (
	"context"
	"log"

	"github.com/97vaibhav/Vessel_tracking/cmd/vesselpb"
	"github.com/97vaibhav/Vessel_tracking/internal/domain/model"
	"github.com/97vaibhav/Vessel_tracking/internal/domain/service"
	"github.com/97vaibhav/Vessel_tracking/internal/errors"
)

func NewVesselServer(vesselService service.VesselUsecase) *server {
	vesselServer := &server{
		service: vesselService,
	}
	return vesselServer
}

type server struct {
	service service.VesselUsecase
	vesselpb.UnimplementedVesselServiceServer
}

func (s *server) CreateVessel(ctx context.Context, in *vesselpb.CreateVesselRequest) (*vesselpb.CreateVesselResponse, error) {
	td := s.ConvertIntoVesselModel(in.Vessel)
	data, err := s.service.CreateVessel(td)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	res := s.ConvertIntoVesselRPC(data)

	return &vesselpb.CreateVesselResponse{
		Id: res.Id,
	}, nil
}

func (s *server) GetVessels(ctx context.Context, in *vesselpb.GetVesselsRequest) (*vesselpb.GetVesselsResponse, error) {
	vessels, err := s.service.GetVessels()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	response := &vesselpb.GetVesselsResponse{
		Vessels: make([]*vesselpb.Vessel, 0),
	}

	for _, vessel := range vessels {
		response.Vessels = append(response.Vessels, s.ConvertIntoVesselRPC(vessel))
	}

	return response, nil
}

func (s *server) UpdateVessel(c context.Context, in *vesselpb.UpdateVesselRequest) (*vesselpb.UpdateVesselResponse, error) {
	vessel := s.ConvertIntoVesselModel(in.Vessel)
	res, err := s.service.UpdateVessel(vessel)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &vesselpb.UpdateVesselResponse{Vessel: s.ConvertIntoVesselRPC(res)}, nil
}

func (s server) GetVessel(ctx context.Context, in *vesselpb.GetVesselRequest) (*vesselpb.GetVesselResponse, error) {
	if in == nil {
		return nil, errors.ErrInvalid
	}
	id := int64(in.Id)
	td, err := s.service.GetVessel(model.VesselID(id))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if td == nil {
		return nil, errors.ErrNotFound
	}
	return &vesselpb.GetVesselResponse{
		Vessel: s.ConvertIntoVesselRPC(td),
	}, nil
}

func (s *server) CreateVoyage(ctx context.Context, in *vesselpb.CreateVoyageRequest) (*vesselpb.CreateVoyageResponse, error) {
	voyageModel := s.ConvertIntoVoyageModel(in.Voyage)

	data, err := s.service.CreateVoyage(voyageModel)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	response := s.ConvertIntoVoyageRPC(data)

	return &vesselpb.CreateVoyageResponse{
		Id: response.Id,
	}, nil
}

func (s *server) UpdateVoyage(c context.Context, in *vesselpb.UpdateVoyageRequest) (*vesselpb.UpdateVoyageResponse, error) {
	voyage := s.ConvertIntoVoyageModel(in.Voyage)
	res, err := s.service.UpdateVoyage(voyage)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &vesselpb.UpdateVoyageResponse{Voyage: s.ConvertIntoVoyageRPC(res)}, nil
}

func (s server) GetVoyage(ctx context.Context, in *vesselpb.GetVoyageRequest) (*vesselpb.GetVoyageResponse, error) {
	if in == nil {
		return nil, errors.ErrInvalid
	}
	voyageId := int64(in.VoyageId)
	td, err := s.service.GetVoyage(model.VoyageID(voyageId), model.VesselID(in.VesselId))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if td == nil {
		return nil, errors.ErrNotFound
	}
	return &vesselpb.GetVoyageResponse{
		Voyage: s.ConvertIntoVoyageRPC(td),
	}, nil
}
