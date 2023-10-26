package datastore

import (
	"fmt"

	"github.com/97vaibhav/Vessel_tracking/internal/domain/model"
	"github.com/97vaibhav/Vessel_tracking/internal/errors"
)

func (m *mysqlVesselRepository) Fetch(query string, args ...interface{}) ([]*model.Vessel, error) {
	rows, err := m.db.Query(query, args...)
	if err != nil {
		fmt.Print(err)
		return nil, errors.ErrInternalServer
	}
	defer rows.Close()
	res := make([]*model.Vessel, 0)
	for rows.Next() {
		t := new(model.Vessel)
		err = rows.Scan(
			&t.Id,
			&t.Name,
			&t.OwnerID,
			&t.NaccsCode,
		)

		if err != nil {
			fmt.Print(err)
			return nil, errors.ErrInternalServer
		}
		res = append(res, t)
	}

	return res, nil
}

func (m *mysqlVesselRepository) FetchVoyage(query string, args ...interface{}) ([]*model.Voyage, error) {
	rows, err := m.db.Query(query, args...)
	if err != nil {
		fmt.Print(err)
		return nil, errors.ErrInternalServer
	}
	defer rows.Close()
	res := make([]*model.Voyage, 0)
	for rows.Next() {
		t := new(model.Voyage)
		err = rows.Scan(
			&t.Id,
			&t.VesselID,
			&t.DepartureLocation,
			&t.ArrivalLocation,
			&t.DepartureTime,
			&t.ArrivalTime,
			&t.Details,
		)

		if err != nil {
			fmt.Print(err)
			return nil, errors.ErrInternalServer
		}
		res = append(res, t)
	}

	return res, nil
}
