package service

import (
	"github.com/97vaibhav/Vessel_tracking/internal/domain/model"
	"github.com/97vaibhav/Vessel_tracking/internal/errors"
	"github.com/97vaibhav/Vessel_tracking/internal/infrastructure/datastore"
)

type vesselService struct {
	vesselService datastore.VesselRepository
}

func NewVesselService(v datastore.VesselRepository) VesselUsecase {
	return &vesselService{v}
}

func (v *vesselService) CreateVessel(m *model.Vessel) (*model.Vessel, error) {
	id, err := v.vesselService.CreateVessel(m)
	if err != nil {
		return nil, err
	}

	m.Id = model.VesselID(id)
	return m, nil
}

func (s *vesselService) GetVessels() ([]*model.Vessel, error) {
	vessels, err := s.vesselService.GetVessels()
	if err != nil {
		return nil, err
	}

	return vessels, nil
}

func (s *vesselService) GetVessel(id model.VesselID) (*model.Vessel, error) {
	vessel, err := s.vesselService.GetVessel(id)
	if err != nil {
		return nil, err
	}

	return vessel, nil
}

func (v *vesselService) UpdateVessel(vessel *model.Vessel) (*model.Vessel, error) {
	_, err := v.vesselService.GetVessel(vessel.Id)
	if err != nil {
		return nil, err
	}
	return v.vesselService.UpdateVessel(vessel)
}

func (v *vesselService) CreateVoyage(m *model.Voyage) (*model.Voyage, error) {
	// First checking if there is any vessel presnt or not
	_, err := v.vesselService.GetVessel(m.VesselID)
	if err != nil {
		return nil, errors.ErrNoVessel
	}
	// letting user to create Voyage only if departure time is less than arrival mtime
	if m.DepartureTime.After(m.ArrivalTime) {
		return nil, errors.ErrInvalidVoyage
	}
	id, err := v.vesselService.CreateVoyage(m)
	if err != nil {
		return nil, err
	}
	m.Id = model.VoyageID(id)
	return m, nil
}

func (v *vesselService) UpdateVoyage(voyage *model.Voyage) (*model.Voyage, error) {

	_, err := v.vesselService.GetVoyage(voyage.Id, voyage.VesselID)
	if err != nil {
		return nil, errors.ErrNoVessel
	}
	return v.vesselService.UpdateVoyage(voyage)
}

func (v vesselService) GetVoyage(id model.VoyageID, vesselID model.VesselID) (*model.Voyage, error) {

	//checking if vessel is present or not ,Since voyage can only be present if and only if there is any vessel
	_, err := v.vesselService.GetVessel(vesselID)
	if err != nil {
		return nil, errors.ErrNoVessel
	}

	voyage, err := v.vesselService.GetVoyage(id, vesselID)
	if err != nil {
		return nil, err
	}
	return voyage, nil
}
