package irkit

func Init() []Device {
	devices := []Device{}

	light := Device{
		ID:         1,
		Name:       "light",
		Status:     false,
		OnCommand:  HomeLightOn,
		OffCommand: HomeLightOff,
	}
	devices = append(devices, light)

	aircon := Device{
		ID:         2,
		Name:       "aircon",
		Status:     false,
		OnCommand:  AirConOn,
		OffCommand: AirConOff,
	}
	devices = append(devices, aircon)

	tv := Device{
		ID:         3,
		Name:       "tv",
		Status:     false,
		OnCommand:  TVPowerToggle,
		OffCommand: TVPowerToggle,
	}
	devices = append(devices, tv)

	return devices
}
