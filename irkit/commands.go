package irkit

import (
	"fmt"
	"time"

	"github.com/whywaita/yayoi/vars"
)

func AirConOn() error {
	err := SendSignal(vars.JsonAirConOn)
	if err != nil {
		err = fmt.Errorf("AirCon On error! : %s", err)
		return err
	}
	return nil
}

func AirConOff() error {
	err := SendSignal(vars.JsonAirConOff)
	if err != nil {
		err = fmt.Errorf("AirCon Off error! : %s", err)
		return err

	}
	return nil
}

func HomeLightOn() error {
	err := SendSignal(vars.JsonHomeLight)
	if err != nil {
		err = fmt.Errorf("HomeLight On error! : %s", err)
		return err
	}
	return nil
}
func HomeLightOff() error {
	for i := 0; i < 4; i++ {
		err := SendSignal(vars.JsonHomeLight)
		if err != nil {
			err = fmt.Errorf("HomeLight Off error! : %s", err)
			return err
		}
		time.Sleep(2 * time.Second)
	}

	return nil
}

func TVPowerToggle() error {
	err := SendSignal(vars.JsonTVPower)
	if err != nil {
		err = fmt.Errorf("TV Power toggle error! : %s", err)
		return err
	}

	return nil
}
