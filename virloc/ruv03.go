package virloc

import (
	"fmt"
	"strconv"
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
	StateDigitalInputs string
	DInput0            string
	DInput1            string
	DInput2            string
	DInput3            string
	DInput4            string
	DInput5            string
	MainSupplyInput    string
	IgnitionInput      string
}

// ToRawMessage implements VirlocReport.
func (r *RUV03) ToRawMessage() string {
	return r.RawData
}

// serialize implements VirlocReport.
func (r *RUV03) serialize(msg string) (VirlocReport, error) {
	msgwr := removeSpecialCharsAndSpaces(msg)
	arrmsg := strings.Split(msgwr, ",")

	var (
		date      string
		time      string
		dgtsinput string
	)

	if _, err := fmt.Sscanf(arrmsg[0], "%5s%3s", &r.PackageType, &r.EventIndexDispatch); err != nil {
		return nil, ErrReadingMessage(err)
	}

	r.ProtocolIdentifier = arrmsg[1]

	if _, err := fmt.Sscanf(arrmsg[2], "%6s%6s", &date, &time); err != nil {
		return nil, ErrReadingMessage(err)
	}

	r.Date = formatDate(date)
	r.Time = formatTime(time)
	dgtsinput = strings.TrimSpace(arrmsg[3])
	r.ThrottlePosition = arrmsg[4]
	r.Hourmeter = arrmsg[5]
	r.Odometer = arrmsg[6]
	r.EngineRotation = arrmsg[7]
	r.EngineTemperature = arrmsg[8]
	r.EnginePressure = arrmsg[9]
	r.FuelLevel = arrmsg[10]
	r.FuelConsumption = arrmsg[11]
	r.Empty1 = arrmsg[12]
	r.Speed = arrmsg[13]
	r.EngineTorque = arrmsg[14]
	r.Empty2 = arrmsg[15]
	r.EngineBrake = arrmsg[16]
	r.Empty3 = arrmsg[17]
	r.Empty4 = arrmsg[18]
	r.Empty5 = arrmsg[19]
	r.CruiseControlState = getonoff(arrmsg[20], "0", "1")
	r.EmploymentState = getonoff(arrmsg[21], "0", "64")
	r.ParkingBrakeState = getonoff(arrmsg[22], "0", "4")
	r.ServiceBrakeState = removeDeviceData(getonoff(arrmsg[23], "0", "8"))

	r.setDigitalInputsState(dgtsinput)

	return r, nil
}

func newRUV03(ms message) VirlocReport {
	return &RUV03{
		message: ms,
	}
}

func (r *RUV03) setDigitalInputsState(dgtinputs string) error {
	sdi, err := strconv.ParseUint(dgtinputs, 16, 32)

	if err != nil {
		return ErrReadingMessage(err)
	}
	r.StateDigitalInputs = fmt.Sprintf("%8b", sdi)

	states := asBits(sdi)
	r.IgnitionInput = getonoff(fmt.Sprint(states[0]), "0", "1")
	r.MainSupplyInput = getonoff(fmt.Sprint(states[1]), "0", "1")
	r.DInput5 = getonoff(fmt.Sprint(states[2]), "0", "1")
	r.DInput4 = getonoff(fmt.Sprint(states[3]), "0", "1")
	r.DInput3 = getonoff(fmt.Sprint(states[4]), "0", "1")
	r.DInput2 = getonoff(fmt.Sprint(states[5]), "0", "1")
	r.DInput1 = getonoff(fmt.Sprint(states[6]), "0", "1")
	r.DInput0 = getonoff(fmt.Sprint(states[7]), "0", "1")

	return nil
}
