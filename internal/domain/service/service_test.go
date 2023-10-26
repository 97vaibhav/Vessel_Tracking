package service_test

import (
	"testing"

	"github.com/97vaibhav/Vessel_tracking/internal/domain/model"
	"github.com/97vaibhav/Vessel_tracking/internal/domain/service"
	"github.com/97vaibhav/Vessel_tracking/mocks"
	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateVessel(t *testing.T) {
	mockVesselRepo := new(mocks.VesselRepository)
	var mockVessel model.Vessel
	err := faker.FakeData(&mockVessel)
	assert.NoError(t, err)
	mockVesselRepo.On("CreateVessel", &mockVessel).Return(int64(mockVessel.Id), nil)
	srv := service.NewVesselService(mockVesselRepo)
	_, err = srv.CreateVessel(&mockVessel)
	assert.NoError(t, err)
}

func TestGetVessel(t *testing.T) {
	mockVesselRepo := new(mocks.VesselRepository)
	var mockVessel model.Vessel
	err := faker.FakeData(&mockVessel)
	assert.NoError(t, err)

	mockVesselRepo.On("GetVessel", mock.AnythingOfType("model.VesselID")).Return(&mockVessel, nil)

	srv := service.NewVesselService(mockVesselRepo)

	td, err := srv.GetVessel(mockVessel.Id)
	assert.NotNil(t, td)
	assert.NoError(t, err)
	assert.Equal(t, mockVessel.Id, td.Id)
	assert.Equal(t, mockVessel.Name, td.Name)
	assert.Equal(t, mockVessel.NaccsCode, td.NaccsCode)
}

func TestUpdateVessel(t *testing.T) {
	mockVesselRepo := new(mocks.VesselRepository)

	var mockVessel model.Vessel
	err := faker.FakeData(&mockVessel)
	assert.NoError(t, err)

	updateVessel := model.Vessel{
		Id:        mockVessel.Id,
		Name:      "name",
		NaccsCode: "code",
		OwnerID:   "updatedowner",
	}

	mockVesselRepo.On("GetVessel", mock.AnythingOfType("model.VesselID")).Return(&mockVessel, nil)
	mockVesselRepo.On("UpdateVessel", &mockVessel).Return(&updateVessel, nil)

	srv := service.NewVesselService(mockVesselRepo)

	td, err := srv.UpdateVessel(&mockVessel)

	assert.NoError(t, err)
	assert.Equal(t, td.Id, updateVessel.Id)
	assert.Equal(t, td.Name, updateVessel.Name)
	assert.Equal(t, td.OwnerID, updateVessel.OwnerID)
}
