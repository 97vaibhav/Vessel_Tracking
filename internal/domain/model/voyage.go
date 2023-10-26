package model

import "time"

type VoyageID int32

type Voyage struct {
	Id                VoyageID
	VesselID          VesselID
	DepartureLocation string
	ArrivalLocation   string
	DepartureTime     time.Time
	ArrivalTime       time.Time
	Details           string
}
