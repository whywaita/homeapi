package main

import "github.com/whywaita/yayoi/irkit"

func Init() []irkit.Device {
	devices := []irkit.Device{}

	light := irkit.Device{
		ID:         1,
		Name:       "light",
		Status:     false,
		OnCommand:  irkit.HomeLightOn,
		OffCommand: irkit.HomeLightOff,
	}
	devices = append(devices, light)

	aircon := irkit.Device{
		ID:         2,
		Name:       "aircon",
		Status:     false,
		OnCommand:  irkit.AirConOn,
		OffCommand: irkit.AirConOff,
	}
	devices = append(devices, aircon)

	tv := irkit.Device{
		ID:         3,
		Name:       "tv",
		Status:     false,
		OnCommand:  irkit.TVPowerToggle,
		OffCommand: irkit.TVPowerToggle,
	}
	devices = append(devices, tv)

	return devices
}
