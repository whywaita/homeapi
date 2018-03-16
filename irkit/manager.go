package irkit

type Manager struct {
	Devices []Device `json:"devices"`
}

func NewManager() Manager {
	manager := Manager{
		Devices: MakeDeviceList(),
	}

	return manager
}

func (m *Manager) SwitchDeviceStatus(index int, sw string) error {
	var err error

	if sw == "on" {
		err = m.Devices[index].OnCommand()
		m.Devices[index].Status = true
	} else {
		err = m.Devices[index].OffCommand()
		m.Devices[index].Status = false
	}

	if err != nil {
		return err
	}

	return nil
}
