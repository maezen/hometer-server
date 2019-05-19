package store

import (
	"hometer-server/model"
	"strconv"
	"strings"
	"time"

	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/conn/physic"
	"periph.io/x/periph/devices/bmxx80"
)

var Temperatures []model.Temperature = []model.Temperature{}
var Humidities []model.Humidity = []model.Humidity{}
var Pressures []model.Pressure = []model.Pressure{}

func StartLoggingValues() {

	for {
		// TODO: Find better solution to cut slices, maybe move last values to front
		if len(Temperatures) >= 1000 {
			Temperatures = Temperatures[:0]
		}

		if len(Humidities) >= 1000 {
			Humidities = Humidities[:0]
		}

		if len(Pressures) >= 1000 {
			Pressures = Pressures[:0]
		}

		// Open a handle to the first available I²C bus:
		bus, err := i2creg.Open("")
		if err != nil {
			continue
		}
		defer bus.Close()

		// Open a handle to a bme280/bmp280 connected on the I²C bus using default
		// settings:
		dev, err := bmxx80.NewI2C(bus, 0x76, &bmxx80.DefaultOpts)
		if err != nil {
			continue
		}
		defer dev.Halt()

		// Read temperature from the sensor:
		var env physic.Env
		if err = dev.Sense(&env); err != nil {
			continue
		}

		humidityString := strings.Replace(env.Humidity.String(), "%rH", "", 1)
		humidityFloat, err := strconv.ParseFloat(humidityString, 64)
		if err != nil {
			continue
		}

		pressureString := strings.Replace(env.Pressure.String(), "kPa", "", 1)
		pressureFloat, err := strconv.ParseFloat(pressureString, 64)
		if err != nil {
			continue
		}

		temperature := model.Temperature{Date: time.Now(), Value: env.Temperature.Celsius()}
		humdidity := model.Humidity{Date: time.Now(), Value: humidityFloat}
		pressure := model.Pressure{Date: time.Now(), Value: pressureFloat}

		Temperatures = append(Temperatures, temperature)
		Humidities = append(Humidities, humdidity)
		Pressures = append(Pressures, pressure)

		// Sleep 5s minute till next value logging
		time.Sleep(5 * time.Minute)
	}
}
