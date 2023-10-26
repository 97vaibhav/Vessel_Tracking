package datastore

import "github.com/97vaibhav/Vessel_tracking/internal/domain/model"

type VesselRepository interface {
	//Interface for Vessels

	CreateVessel(vessel *model.Vessel) (int64, error)
	GetVessels() ([]*model.Vessel, error)
	UpdateVessel(vessel *model.Vessel) (*model.Vessel, error)
	GetVessel(model.VesselID) (*model.Vessel, error)

	// Interface for Voyages

	CreateVoyage(voyage *model.Voyage) (int64, error)
	UpdateVoyage(voyage *model.Voyage) (*model.Voyage, error)
	GetVoyage(model.VoyageID, model.VesselID) (*model.Voyage, error)
}
