package virloc

import (
	"fmt"
	"strings"
)

type RUV02 struct {
	message
	PackageType              string
	EventIndexDispatch       string
	ProtocolIdentifier       string
	Date                     string
	Time                     string
	RPMRange1                string
	RPMRange2                string
	RPMRange3                string
	RPMRange4                string
	RPMRange5                string
	TimeInInertia            string
	EngineBrakingTime        string
	TractionFreeMovementTime string
	TimeOfTravel             string
	TravelledDistance        string
	FuelConsumed             string
}

// ToRawMessage implements VirlocReport.
func (r *RUV02) ToRawMessage() string {
	return r.RawData
}

// serialize implements VirlocReport.
func (r *RUV02) serialize(msg string) (VirlocReport, error) {
	msgr := removeSpecialCharsAndSpaces(msg)

	msgarr := strings.Split(msgr, ",")

	if _, err := fmt.Sscanf(msgarr[0], "%5s%3s", &r.PackageType, &r.EventIndexDispatch); err != nil {
		return nil, ErrReadingMessage(err)
	}

	r.ProtocolIdentifier = msgarr[1]

	if _, err := fmt.Sscanf(msgarr[2], "%6s%6s", &r.Date, &r.Time); err != nil {
		return nil, ErrReadingMessage(err)
	}

	r.RPMRange1 = msgarr[3]
	r.RPMRange2 = msgarr[4]
	r.RPMRange3 = msgarr[5]
	r.RPMRange4 = msgarr[6]
	r.RPMRange5 = msgarr[7]
	r.TimeInInertia = msgarr[8]
	r.EngineBrakingTime = msgarr[9]
	r.TractionFreeMovementTime = msgarr[10]
	r.TimeOfTravel = msgarr[11]
	r.TravelledDistance = msgarr[12]
	r.FuelConsumed = removeDeviceData(msgarr[13])

	return r, nil

}

func newRUV02(ms message) VirlocReport {
	return &RUV02{
		message: ms,
	}
}
