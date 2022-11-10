package repository

import (
	"fmt"
	"truphone-api/src/errs"
	"truphone-api/src/models"
)

func (db DB) CreateDevice(device models.Device) error {
	sql := db.sql
	query := `INSERT INTO public.device ("name", brand, creation_time) VALUES($1, $2, $3);`

	_, err := sql.Exec(query, device.Name, device.Brand, device.CreationTime)
	if err != nil {
		return fmt.Errorf("%s %w", errs.ERR_EXEC_QUERY, err)
	}

	defer sql.Close()
	return nil
}

func (db DB) GetllAllDevices() ([]models.Device, error) {
	sql := db.sql
	query := `SELECT id, "name", brand, creation_time FROM public.device;`

	rows, err := sql.Query(query)
	if err != nil {
		return nil, fmt.Errorf("%s %w", errs.ERR_EXEC_QUERY, err)
	}

	deviceList := make([]models.Device, 0)
	for rows.Next() {
		var dv models.Device
		if err := rows.Scan(
			&dv.ID,
			&dv.Name,
			&dv.Brand,
			&dv.CreationTime,
		); err != nil {
			return nil, err
		}
		deviceList = append(deviceList, dv)
	}

	return deviceList, nil
}

func (db DB) GetDevicesByID(id string) (*models.Device, error) {
	sql := db.sql
	query := `SELECT id, "name", brand, creation_time FROM public.device WHERE id=$1;`

	rows, err := sql.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("%s %w", errs.ERR_EXEC_QUERY, err)
	}

	dv := new(models.Device)
	for rows.Next() {
		if err := rows.Scan(
			&dv.ID,
			&dv.Name,
			&dv.Brand,
			&dv.CreationTime,
		); err != nil {
			return nil, err
		}
	}

	return dv, nil
}

func (db DB) UpdateDevice(device models.Device) error {
	sql := db.sql
	var args string
	query := `UPDATE public.device SET `
	if len(device.Name) > 0 {
		args = fmt.Sprintf(`name='%s',`, device.Name)
		query += args
	}

	if len(device.Brand) > 0 {
		args = fmt.Sprintf(`brand=%v,`, device.Brand)
		query += args
	}

	if len(device.CreationTime) > 0 {
		args = fmt.Sprintf(`creation_time=%v,`, device.CreationTime)
		query += args
	}

	query += ` WHERE id=$1`

	_, err := sql.Exec(query)
	if err != nil {
		return fmt.Errorf("%s %w", errs.ERR_EXEC_QUERY, err)
	}

	return nil
}

func (db DB) DeleteDevice(id string) error {
	sql := db.sql
	query := `DELETE FROM public.device WHERE id=$1;`

	_, err := sql.Exec(query, id)
	if err != nil {
		return fmt.Errorf("%s %w", errs.ERR_EXEC_QUERY, err)
	}

	return nil
}

func (db DB) GetDevicesByBrand(brand string) (*models.Device, error) {
	sql := db.sql
	query := `SELECT id, "name", brand, creation_time FROM public.device WHERE brand=$1;`

	rows, err := sql.Query(query, brand)
	if err != nil {
		return nil, fmt.Errorf("%s %w", errs.ERR_EXEC_QUERY, err)
	}

	dv := new(models.Device)
	for rows.Next() {
		if err := rows.Scan(
			&dv.ID,
			&dv.Name,
			&dv.Brand,
			&dv.CreationTime,
		); err != nil {
			return nil, err
		}
	}

	return dv, nil

}
