package irkit

type Device struct {
	ID         int          `json:"id"`
	Name       string       `json:"name"`
	Alias      string       `json:"alias"`
	Status     bool         `json:"status"`
	DescStatus string       `json:"desc_status"`
	OnCommand  func() error `json:"on_command"`
	OffCommand func() error `json:"off_command"`
}

type UserDevice struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Status     bool   `json:"status"`
	DescStatus string `json:"desc_status"`
}

type UserDeviceList struct {
	Devices []UserDevice `json:"devices"`
}

func (d *Device) ToUserDevice() UserDevice {
	u := UserDevice{}

	u.ID = d.ID
	u.Name = d.Name
	u.Status = d.Status
	u.DescStatus = d.DescStatus

	return u
}
