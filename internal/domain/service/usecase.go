package service

import "github.com/97vaibhav/Vessel_tracking/internal/domain/model"

type VesselUsecase interface {
	// Interface for Vessels
	CreateVessel(vessel *model.Vessel) (*model.Vessel, error)
	UpdateVessel(vessel *model.Vessel) (*model.Vessel, error)
	GetVessels() ([]*model.Vessel, error)
	GetVessel(model.VesselID) (*model.Vessel, error)

	//Interface for Voyages
	CreateVoyage(voyage *model.Voyage) (*model.Voyage, error)
	UpdateVoyage(voyage *model.Voyage) (*model.Voyage, error)
	GetVoyage(voyageId model.VoyageID, vesselId model.VesselID) (*model.Voyage, error)
}
