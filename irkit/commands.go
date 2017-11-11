package irkit

import (
	"time"

	"github.com/whywaita/homeapi/vars"

	"go.uber.org/zap"
)

func AirConOn(logger *zap.Logger) {
	err := SendSignal(vars.JsonAirConOn)
	if err != nil {
		logger.Warn("AirCon On error! :", zap.Error(err))
	}
	return
}

func AirConOff(logger *zap.Logger) {
	err := SendSignal(vars.JsonAirConOff)
	if err != nil {
		logger.Warn("AirCon On error! :", zap.Error(err))
	}
	return
}

func HomeLightOn(logger *zap.Logger) {
	err := SendSignal(vars.JsonHomeLight)
	if err != nil {
		logger.Warn("HomeLight On error! :", zap.Error(err))
	}
	return
}
func HomeLightOff(logger *zap.Logger) {
	for i := 0; i < 4; i++ {
		err := SendSignal(vars.JsonHomeLight)
		if err != nil {
			logger.Warn("HomeLight Off error! :", zap.Error(err))
			time.Sleep(2 * time.Second)
		}
	}
}
