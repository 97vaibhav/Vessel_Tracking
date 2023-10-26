package handler

import (
	"github.com/97vaibhav/Vessel_tracking/cmd/vesselpb"
	"github.com/97vaibhav/Vessel_tracking/internal/domain/model"
	"github.com/golang/protobuf/ptypes"
)

func (s *server) ConvertIntoVesselModel(in *vesselpb.Vessel) *model.Vessel {
	return &model.Vessel{
		Id:        model.VesselID(in.Id),
		Name:      in.Name,
		OwnerID:   in.OwnerId,
		NaccsCode: in.NaccsCode,
	}
}

func (s *server) ConvertIntoVesselRPC(in *model.Vessel) *vesselpb.Vessel {
	if in == nil {
		return nil
	}
	return &vesselpb.Vessel{
		Id:        int32(in.Id),
		Name:      in.Name,
		OwnerId:   in.OwnerID,
		NaccsCode: in.NaccsCode,
	}
}

func (s *server) ConvertIntoVoyageModel(in *vesselpb.Voyage) *model.Voyage {
	departureTime, _ := ptypes.Timestamp(in.DepartureTime)
	arrivalTime, _ := ptypes.Timestamp(in.ArrivalTime)

	return &model.Voyage{
		// Map fields from the gRPC message to the model.
		Id:                model.VoyageID(in.Id),
		VesselID:          model.VesselID(in.VesselId),
		DepartureLocation: in.DepartureLocation,
		ArrivalLocation:   in.ArrivalLocation,
		DepartureTime:     departureTime,
		ArrivalTime:       arrivalTime,
		Details:           in.Details,
	}
}

func (s *server) ConvertIntoVoyageRPC(in *model.Voyage) *vesselpb.Voyage {
	departureTime, _ := ptypes.TimestampProto(in.DepartureTime)
	arrivalTime, _ := ptypes.TimestampProto(in.ArrivalTime)

	if in == nil {
		return nil
	}
	return &vesselpb.Voyage{
		Id:                int32(in.Id),
		VesselId:          int32(in.VesselID),
		DepartureLocation: in.DepartureLocation,
		ArrivalLocation:   in.ArrivalLocation,
		DepartureTime:     departureTime,
		ArrivalTime:       arrivalTime,
		Details:           in.Details,
	}
}
