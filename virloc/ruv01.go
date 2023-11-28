package virloc

import (
	"fmt"
	"strings"
)

type RUV01 struct {
	message
	PackageType                 string
	EventIndexDispatch          string
	ProtocolIdentifier          string
	Date                        string
	Time                        string
	Latitude                    string
	Longitude                   string
	Speed                       string
	Direction                   string
	GpsState                    string
	SecondsSinceLastPosition    string
	StateDigitalInputs          string
	Empty1                      string
	HDOPLastPosition            string
	BackupBatteryVoltage        string
	MainSupplyVoltage           string
	MaxSpeedOnViolationEvent    string
	MaxRotationOnViolationEvent string
	Empty2                      string
	Hourmeter                   string
	Odometer                    string
	EngineRotation              string
	EngineTemperature           string
	EngineOilPressure           string
	FuelLevel                   string
	SpeedEventInRain            string
	ConnectionStatusOnEvent     string
	Connection4G                string
	DriverId                    string
}

func (r *RUV01) ToRawMessage() string {
	return r.RawData
}

func (r *RUV01) serialize(msg string) (VirlocReport, error) {
	message := removeSpecialCharsAndSpaces(msg)

	arrmes := strings.Split(message, ",")
	_, err := fmt.Sscanf(arrmes[0], "%5s%3s", &r.PackageType, &r.EventIndexDispatch)

	if err != nil {
		return nil, ErrReadingMessage(err)
	}

	r.ProtocolIdentifier = arrmes[1]

	_, err = fmt.Sscanf(arrmes[2], "%6s%6s%8s%9s%3s%3s%1s%2s%2s%2s%2s", &r.Date,
		&r.Time,
		&r.Latitude,
		&r.Longitude,
		&r.Speed,
		&r.Direction,
		&r.GpsState,
		&r.SecondsSinceLastPosition,
		&r.StateDigitalInputs,
		&r.Empty1,
		&r.HDOPLastPosition)

	if err != nil {
		return nil, ErrReadingMessage(err)
	}

	_, err = fmt.Sscanf(arrmes[3], "%4s%4s", &r.BackupBatteryVoltage, &r.MainSupplyVoltage)

	if err != nil {
		return nil, ErrReadingMessage(err)
	}

	r.MaxSpeedOnViolationEvent = arrmes[4]
	r.MaxRotationOnViolationEvent = arrmes[5]
	r.Empty2 = arrmes[6]
	r.Hourmeter = arrmes[7]
	r.Odometer = arrmes[8]
	r.EngineRotation = arrmes[9]
	r.EngineTemperature = arrmes[10]
	r.EngineOilPressure = arrmes[11]
	r.FuelLevel = arrmes[12]
	r.SpeedEventInRain = arrmes[13]
	r.ConnectionStatusOnEvent = arrmes[14]
	r.Connection4G = arrmes[15]
	r.DriverId = removeDeviceData(arrmes[16])

	return r, nil

}

func newRUV01(ms message) VirlocReport {
	r := &RUV01{
		message: ms,
	}

	return r
}
