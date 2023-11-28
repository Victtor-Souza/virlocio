package virloc

import (
	"fmt"
	"strings"
)

type RUV03 struct {
	message
	PackageType        string
	EventIndexDispatch string
	ProtocolIdentifier string
	Date               string
	Time               string
	ThrottlePosition   string
	Hourmeter          string
	Odometer           string
	EngineRotation     string
	EngineTemperature  string
	EnginePressure     string
	FuelLevel          string
	FuelConsumption    string
	Empty1             string
	Speed              string
	EngineTorque       string
	Empty2             string
	EngineBrake        string
	Empty3             string
	Empty4             string
	Empty5             string
	CruiseControlState string
	EmploymentState    string
	ParkingBrakeState  string
	ServiceBrakeState  string
}

// ToRawMessage implements VirlocReport.
func (r *RUV03) ToRawMessage() string {
	return r.RawData
}

// serialize implements VirlocReport.
func (r *RUV03) serialize(msg string) (VirlocReport, error) {
	msgwr := removeSpecialCharsAndSpaces(msg)
	arrmsg := strings.Split(msgwr, ",")

	if _, err := fmt.Sscanf(arrmsg[0], "%5s%3s", &r.PackageType, &r.EventIndexDispatch); err != nil {
		return nil, ErrReadingMessage(err)
	}

	r.ProtocolIdentifier = arrmsg[1]

	if _, err := fmt.Sscanf(arrmsg[2], "%6s%6s", &r.Date, &r.Time); err != nil {
		return nil, ErrReadingMessage(err)
	}

	r.ThrottlePosition = arrmsg[3]
	r.Hourmeter = arrmsg[4]
	r.Odometer = arrmsg[5]
	r.EngineRotation = arrmsg[6]
	r.EngineTemperature = arrmsg[7]
	r.EnginePressure = arrmsg[8]
	r.FuelLevel = arrmsg[9]
	r.FuelConsumption = arrmsg[10]
	r.Empty1 = arrmsg[11]
	r.Speed = arrmsg[12]
	r.EngineTorque = arrmsg[13]
	r.Empty2 = arrmsg[14]
	r.EngineBrake = arrmsg[15]
	r.Empty3 = arrmsg[16]
	r.Empty4 = arrmsg[17]
	r.Empty5 = arrmsg[18]
	r.CruiseControlState = arrmsg[19]
	r.EmploymentState = arrmsg[20]
	r.ParkingBrakeState = arrmsg[21]
	r.ServiceBrakeState = removeDeviceData(arrmsg[22])

	return r, nil
}

func newRUV03(ms message) VirlocReport {
	return &RUV03{
		message: ms,
	}
}
