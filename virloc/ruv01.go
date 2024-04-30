package virloc

import (
	"fmt"
	"strconv"
	"strings"
)

type RUV01 struct {
	message
	PackageType                 string
	EventIndexDispatch          string
	ProtocolIdentifier          string
	Date                        string
	Time                        string
	Latitude                    float64
	Longitude                   float64
	Speed                       string
	Direction                   string
	GpsState                    string
	SecondsSinceLastPosition    string
	StateDigitalInputs          string
	DInput0                     string
	DInput1                     string
	DInput2                     string
	DInput3                     string
	DInput4                     string
	DInput5                     string
	MainSupplyInput             string
	IgnitionInput               string
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

	var (
		date, time, lat, long, dgtinputs, scdslp string
	)

	arrmes := strings.Split(message, ",")
	_, err := fmt.Sscanf(arrmes[0], "%5s%3s", &r.PackageType, &r.EventIndexDispatch)

	if err != nil {
		return nil, ErrReadingMessage(err)
	}

	r.ProtocolIdentifier = arrmes[1]

	_, err = fmt.Sscanf(arrmes[2], "%6s%6s%8s%9s%3s%3s%1s%2s%2s%2s%2s", &date,
		&time,
		&lat,
		&long,
		&r.Speed,
		&r.Direction,
		&r.GpsState,
		&scdslp,
		&dgtinputs,
		&r.Empty1,
		&r.HDOPLastPosition)

	if err != nil {
		return nil, ErrReadingMessage(err)
	}

	_, err = fmt.Sscanf(arrmes[3], "%4s%4s", &r.BackupBatteryVoltage, &r.MainSupplyVoltage)

	if err != nil {
		return nil, ErrReadingMessage(err)
	}

	if err := r.setDigitalInputsState(dgtinputs); err != nil {
		return nil, ErrReadingMessage(err)
	}

	sslp, err := strconv.ParseUint(scdslp, 16, 32)
	if err != nil {
		return nil, ErrReadingMessage(err)
	}

	r.Date = formatDate(date)
	r.Time = formatTime(time)
	r.Latitude = convertStringToFloat64(lat, 2)
	r.Longitude = convertStringToFloat64(long, 3)
	r.SecondsSinceLastPosition = fmt.Sprintf("%v", sslp)
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

func (r *RUV01) setDigitalInputsState(dgtinputs string) error {
	sdi, err := strconv.ParseUint(dgtinputs, 16, 32)

	if err != nil {
		return ErrReadingMessage(err)
	}
	r.StateDigitalInputs = fmt.Sprintf("%8b", sdi)

	states := asBits(sdi)
	r.IgnitionInput = getonoff(fmt.Sprint(states[0]), "0", "1")
	r.MainSupplyInput = getonoff(fmt.Sprint(states[1]), "0", "1")
	r.DInput5 = getonoff(fmt.Sprint(states[2]), "1", "0")
	r.DInput4 = getonoff(fmt.Sprint(states[3]), "1", "0")
	r.DInput3 = getonoff(fmt.Sprint(states[4]), "1", "0")
	r.DInput2 = getonoff(fmt.Sprint(states[5]), "1", "0")
	r.DInput1 = getonoff(fmt.Sprint(states[6]), "1", "0")
	r.DInput0 = getonoff(fmt.Sprint(states[7]), "1", "0")

	return nil
}
