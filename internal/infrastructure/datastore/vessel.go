package datastore

import (
	"database/sql"
	"fmt"

	"github.com/97vaibhav/Vessel_tracking/internal/domain/model"
	"github.com/97vaibhav/Vessel_tracking/internal/errors"
)

type mysqlVesselRepository struct {
	db *sql.DB
}

func NewMysqlVesselRepository(db *sql.DB) VesselRepository {
	return &mysqlVesselRepository{db}
}

// For Vessels related Operations like Create , Update, Get ,GetLists
func (m *mysqlVesselRepository) CreateVessel(a *model.Vessel) (int64, error) {
	const query = "INSERT INTO Vessel (name, owner_id,naccs_code) VALUES (?, ?,?)"
	stmt, err := m.db.Prepare(query)
	if err != nil {
		fmt.Print(err)
		return 0, errors.ErrInternalServer
	}
	res, err := stmt.Exec(a.Name, a.OwnerID, a.NaccsCode)
	if err != nil {
		fmt.Print(err)
		return 0, errors.ErrInternalServer
	}
	return res.LastInsertId()
}

func (m *mysqlVesselRepository) GetVessels() ([]*model.Vessel, error) {
	query := "SELECT id, name, owner_id, naccs_code FROM Vessel"
	rows, err := m.db.Query(query)
	if err != nil {
		return nil, errors.ErrInternalServer
	}
	defer rows.Close()

	vessels := make([]*model.Vessel, 0)
	for rows.Next() {
		var vessel model.Vessel
		if err := rows.Scan(&vessel.Id, &vessel.Name, &vessel.OwnerID, &vessel.NaccsCode); err != nil {
			return nil, errors.ErrInternalServer
		}
		vessels = append(vessels, &vessel)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.ErrInternalServer
	}

	return vessels, nil
}

func (m *mysqlVesselRepository) UpdateVessel(vessel *model.Vessel) (*model.Vessel, error) {
	const query = "UPDATE Vessel SET name=?, owner_id=?, naccs_code=? WHERE id = ?"

	stmt, err := m.db.Prepare(query)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	res, err := stmt.Exec(vessel.Name, vessel.OwnerID, vessel.NaccsCode, vessel.Id)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	if rowsAffected < 1 {
		return nil, errors.ErrNotFound
	}

	return vessel, nil
}

func (m *mysqlVesselRepository) GetVessel(id model.VesselID) (*model.Vessel, error) {
	const query = `SELECT id,name,owner_id,naccs_code FROM Vessel WHERE ID = ?`

	list, err := m.Fetch(query, id)
	if err != nil {
		return nil, errors.ErrNotFound
	}

	if len(list) == 0 {
		return nil, errors.ErrNotFound
	}
	return list[0], nil
}

// For Voyages related operations

func (m *mysqlVesselRepository) CreateVoyage(voyage *model.Voyage) (int64, error) {
	const voyageQuery = "INSERT INTO Voyages (vessel_id, departure_location, arrival_location, departure_time, arrival_time, details) VALUES (?, ?, ?, ?, ?, ?)"
	stmt, err := m.db.Prepare(voyageQuery)
	if err != nil {
		fmt.Print(err)
		return 0, errors.ErrInternalServer
	}
	res, err := stmt.Exec(voyage.VesselID, voyage.DepartureLocation, voyage.ArrivalLocation, voyage.DepartureTime, voyage.ArrivalTime, voyage.Details)
	if err != nil {
		fmt.Print(err)
		return 0, errors.ErrInternalServer
	}
	return res.LastInsertId()
}

func (m *mysqlVesselRepository) GetVoyage(id model.VoyageID, vesselId model.VesselID) (*model.Voyage, error) {
	const query = `SELECT voyage_id,vessel_id,departure_location,arrival_location,departure_time,arrival_time,details FROM Voyages WHERE voyage_id = ?`

	list, err := m.FetchVoyage(query, id)
	if err != nil {
		return nil, errors.ErrNotFound
	}

	if len(list) == 0 {
		return nil, errors.ErrNotFound
	}
	return list[0], nil
}

func (m *mysqlVesselRepository) UpdateVoyage(voyage *model.Voyage) (*model.Voyage, error) {
	const query = "UPDATE Voyages SET departure_location=?, arrival_location=?, departure_time=?, arrival_time=?, details=? WHERE voyage_id = ?"

	stmt, err := m.db.Prepare(query)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	res, err := stmt.Exec(voyage.DepartureLocation, voyage.ArrivalLocation, voyage.DepartureTime, voyage.ArrivalTime, voyage.Details, voyage.Id)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	if rowsAffected < 1 {
		return nil, errors.ErrNotFound
	}

	return voyage, nil
}
