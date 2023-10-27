package errors

import "errors"

var (
	ErrInternalServer = errors.New("internal Server Error")
	ErrNotFound       = errors.New("record Not found")
	ErrInvalid        = errors.New("please enter a valid value ")
	ErrInvalidVoyage  = errors.New("invalid time input because departure time cant be greater than arrival time ")
	ErrNoVessel       = errors.New("no Vessel Present Check Vessel is presnt in database or not ")
)
