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
	Odometer           int64
	EngineRotation     int64
	EngineTemperature  string
	EnginePressure     string
	FuelLevel          int64
	FuelConsumption    string
	Empty1             string
	Speed              int64
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
	Output0            string
	Output1            string
	Output2            string
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
	stsout := arrmsg[4]
	r.ThrottlePosition = arrmsg[5]
	r.Hourmeter = arrmsg[6]
	r.Odometer = convertStringToInt64(arrmsg[7])
	r.EngineRotation = convertStringToInt64(arrmsg[8])
	r.EngineTemperature = arrmsg[9]
	r.EnginePressure = arrmsg[10]
	r.FuelLevel = convertStringToInt64(arrmsg[11])
	r.FuelConsumption = arrmsg[12]
	r.Empty1 = arrmsg[13]
	r.Speed = convertStringToInt64(arrmsg[14])
	r.EngineTorque = arrmsg[15]
	r.Empty2 = arrmsg[16]
	r.EngineBrake = arrmsg[17]
	r.Empty3 = arrmsg[18]
	r.Empty4 = arrmsg[19]
	r.Empty5 = arrmsg[20]
	r.CruiseControlState = getonoff(arrmsg[21], "0", "1")
	r.EmploymentState = getonoff(arrmsg[22], "0", "64")
	r.ParkingBrakeState = getonoff(arrmsg[23], "0", "4")
	r.ServiceBrakeState = removeDeviceData(getonoff(arrmsg[24], "0", "8"))

	r.setDigitalInputsState(dgtsinput)
	r.setOutputsStates(stsout)

	return r, nil
}

func newRUV03(ms message) VirlocReport {
	return &RUV03{
		message: ms,
	}
}

func (r *RUV03) setOutputsStates(stateoutputs string) error {
	sout := strings.Split(stateoutputs, "")

	r.Output0 = getonoff(sout[0], "0", "1")
	r.Output1 = getonoff(sout[1], "0", "1")
	r.Output2 = getonoff(sout[2], "0", "1")

	return nil
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
