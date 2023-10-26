package model

type VesselID int32

type Vessel struct {
	Id        VesselID
	Name      string
	OwnerID   string
	NaccsCode string
}
