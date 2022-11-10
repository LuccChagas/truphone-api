package router

import (
	"net/http"
	"truphone-api/src/errs"
	"truphone-api/src/models"
	"truphone-api/src/services"

	"github.com/labstack/echo/v4"
)

var service services.Store

func AddNewDevice(c echo.Context) error {
	var device models.Device
	if err := c.Bind(&device); err != nil {
		eh := errs.Handling{Error: err.Error(), Message: errs.ERR_BIND_OBJECT}
		return c.JSON(http.StatusBadRequest, eh)
	}

	err := service.CreateDevice(device)
	if err != nil {
		eh := errs.Handling{Error: err.Error(), Message: errs.ERR_SERVICE}
		return c.JSON(http.StatusBadRequest, eh)
	}

	return c.NoContent(http.StatusCreated)
}

func GetAllDevices(c echo.Context) error {

	devices, err := service.GetAllDevices()
	if err != nil {
		eh := errs.Handling{Error: err.Error(), Message: errs.ERR_SERVICE}
		return c.JSON(http.StatusBadRequest, eh)
	}

	return c.JSON(http.StatusOK, devices)
}

func GetDevicesByID(c echo.Context) error {
	id := c.QueryParam("id")
	if len(id) <= 0 {
		return c.JSON(http.StatusBadRequest, "ID is required")
	}

	device, err := service.GetDevicesByID(id)
	if err != nil {
		eh := errs.Handling{Error: err.Error(), Message: errs.ERR_SERVICE}
		return c.JSON(http.StatusBadRequest, eh)
	}

	return c.JSON(http.StatusOK, device)
}

func DeleteDeviceByID(c echo.Context) error {
	id := c.QueryParam("id")
	if len(id) <= 0 {
		return c.JSON(http.StatusBadRequest, "ID is required")
	}

	err := service.DeleteDevice(id)
	if err != nil {
		eh := errs.Handling{Error: err.Error(), Message: errs.ERR_SERVICE}
		return c.JSON(http.StatusBadRequest, eh)
	}

	return c.NoContent(http.StatusNoContent)

}

func GetDeviceByBrand(c echo.Context) error {
	name := c.QueryParam("name")
	if len(name) <= 0 {
		return c.JSON(http.StatusBadRequest, "Name is required")
	}

	device, err := service.GetDeviceByBrand(name)
	if err != nil {
		eh := errs.Handling{Error: err.Error(), Message: errs.ERR_SERVICE}
		return c.JSON(http.StatusBadRequest, eh)
	}

	return c.JSON(http.StatusOK, device)

}

func UpdateDevice(c echo.Context) error {
	var device models.Device

	if err := c.Bind(&device); err != nil {
		eh := errs.Handling{Error: err.Error(), Message: errs.ERR_BIND_OBJECT}
		return c.JSON(http.StatusBadRequest, eh)
	}

	err := service.UpdateDevice(device)
	if err != nil {
		eh := errs.Handling{Error: err.Error(), Message: errs.ERR_SERVICE}
		return c.JSON(http.StatusBadRequest, eh)
	}

	return c.NoContent(http.StatusCreated)
}
