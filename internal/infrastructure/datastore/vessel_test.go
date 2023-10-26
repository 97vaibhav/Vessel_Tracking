package datastore_test

import (
	"errors"
	"testing"

	"github.com/97vaibhav/Vessel_tracking/internal/domain/model"
	"github.com/97vaibhav/Vessel_tracking/internal/infrastructure/datastore"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetVessel(t *testing.T) {
	query := "SELECT id,name,owner_id,naccs_code FROM Vessel WHERE ID = ?"
	t.Run("Record not found", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("OOPS!! an error '%s' was not expected when opening database connection", err)
		}
		t.Cleanup(func() {
			db.Close()
		})

		id := 10
		expectedError := errors.New("record Not found")
		mock.ExpectQuery(query).WithArgs(id).WillReturnError(expectedError)
		repo := datastore.NewMysqlVesselRepository(db)
		vessel, err := repo.GetVessel(model.VesselID(id))
		assert.Nil(t, vessel)
		assert.EqualError(t, err, expectedError.Error(), "error message is 'record not found'")
	})
	t.Run("Successful GetVesssel", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("OOPS!! an error '%s' was not expected when opening database connection", err)
		}
		t.Cleanup(func() {
			db.Close()
		})
		id := int64(1)
		rows := sqlmock.NewRows([]string{"id", "name", "owner_id", "naccs_code"}).
			AddRow(id, "name1", "owner1", "Naccscode")
		mock.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)
		repo := datastore.NewMysqlVesselRepository(db)
		vessel, err := repo.GetVessel(model.VesselID(id))
		assert.NoError(t, err)
		assert.NotNil(t, vessel)
		assert.Equal(t, vessel.Name, "name1")
		assert.Equal(t, vessel.OwnerID, "owner1")
		assert.Equal(t, vessel.NaccsCode, "Naccscode")
	})
}

func TestCreateVessel(t *testing.T) {
	query := "INSERT INTO Vessel"
	vs := &model.Vessel{
		Name:      "test_name",
		OwnerID:   "ownerid",
		NaccsCode: "code",
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("OOPS!! an error '%s' was not expected when opening database connection", err)
	}
	t.Cleanup(func() {
		db.Close()
	})

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(vs.Name, vs.OwnerID, vs.NaccsCode).WillReturnResult(sqlmock.NewResult(99, 1))
	repo := datastore.NewMysqlVesselRepository(db)

	lastID, err := repo.CreateVessel(vs)
	assert.NoError(t, err)
	assert.Equal(t, int64(99), lastID)
}

func TestUpdateVesssel(t *testing.T) {
	query := "UPDATE Vessel SET name=\\?, owner_id=\\?, naccs_code=\\? WHERE id = \\?"

	t.Run("UpdateVessel Success ", func(t *testing.T) {
		vs := &model.Vessel{
			Id:        13,
			Name:      "updated name",
			OwnerID:   "updated owner",
			NaccsCode: "code",
		}

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("OOPS!! an error '%s' was not expected when opening database connection", err)
		}
		t.Cleanup(func() {
			db.Close()
		})

		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(vs.Name, vs.OwnerID, vs.NaccsCode, vs.Id).WillReturnResult(sqlmock.NewResult(13, 1))

		repo := datastore.NewMysqlVesselRepository(db)

		res, err := repo.UpdateVessel(vs)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("UpdateVessel Record Not Found ", func(t *testing.T) {
		vs := &model.Vessel{
			Id:        13,
			Name:      "updated name",
			OwnerID:   "updated owner",
			NaccsCode: "code",
		}

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("OOPS!! an error '%s' was not expected when opening database connection", err)
		}
		t.Cleanup(func() {
			db.Close()
		})

		expectedError := errors.New("record Not found")

		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(vs.Name, vs.OwnerID, vs.NaccsCode, vs.Id).WillReturnResult(sqlmock.NewResult(0, 0))

		repo := datastore.NewMysqlVesselRepository(db)
		res, err := repo.UpdateVessel(vs)
		assert.EqualError(t, err, expectedError.Error(), "error message is 'record not found'")
		assert.Nil(t, res)
	})
}
