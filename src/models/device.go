package models

type Device struct {
	ID           uint32 `json:"id"`
	Name         string `json:"name,omitempty"`
	Brand        string `json:"brand,omitempty"`
	CreationTime string `json:"creation_time"`
}

type DeviceLogicalFields struct {
	IsActive  bool   `json:"isActive,,omitempty"`
	UpdatedAt string `json:"updated_at,,omitempty"`
}
