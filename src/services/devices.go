package services

import (
	"errors"
	"reflect"
	"truphone-api/src/models"
	"truphone-api/src/repository"
)

type Store struct {
	Repo repository.DB
}

func (store Store) CreateDevice(device models.Device) error {

	if reflect.ValueOf(device).IsZero() {
		return errors.New("reflect.ValueOf - Empty Struct")
	}

	err := store.Repo.CreateDevice(device)
	if err != nil {
		return err
	}
	return nil
}

func (store Store) GetAllDevices() ([]models.Device, error) {
	devices, err := store.Repo.GetllAllDevices()
	if err != nil {
		return nil, err
	}

	return devices, nil
}

func (store Store) GetDevicesByID(id string) (*models.Device, error) {
	device, err := store.Repo.GetDevicesByID(id)
	if err != nil {
		return nil, err
	}

	return device, nil
}

func (store Store) DeleteDevice(id string) error {
	err := store.Repo.DeleteDevice(id)
	if err != nil {
		return err
	}

	return nil
}

func (store Store) UpdateDevice(device models.Device) error {

	if reflect.ValueOf(device).IsZero() {
		return errors.New("reflect.ValueOf - Empty Struct")
	}
	err := store.Repo.UpdateDevice(device)
	if err != nil {
		return err
	}

	return nil
}

func (store Store) GetDeviceByBrand(brand string) (*models.Device, error) {

	device, err := store.Repo.GetDevicesByBrand(brand)
	if err != nil {
		return nil, err
	}

	return device, nil
}
