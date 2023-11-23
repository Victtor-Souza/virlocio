package virloc

import (
	"fmt"
)

type QGP struct {
	RawData                     string
	PackageType                 string
	Header                      Header
	DateTimeEvent               string
	Latitude                    string
	Longitude                   string
	Speed                       string
	Direction                   string
	GpsState                    string
	SecondsSinceLastPosition    string
	StateDigitalInputs          string
	EventNumber                 string
	DilutionHorizontalPrecision string
}

func (qgp *QGP) AcceptMessage() string {
	ackmsg := fmt.Sprintf(">ACK;%s;%s", qgp.Header.DeviceId, qgp.Header.MessageNumber)
	chksum := CalculateChecksum(qgp.RawData)
	return fmt.Sprintf("%s;*%s<", ackmsg, chksum)
}

func (qgp *QGP) ToRawMessage() string {
	return qgp.RawData
}

func (qgp *QGP) serialize(msg string) VirlocReport {
	messagewspace := RemoveSpecialCharsAndSpaces(msg)
	_, err := fmt.Sscanf(messagewspace, "%3s%12s%8s%9s%3s%3s%1s%2s%2s%2s%2s",
		&qgp.PackageType,
		&qgp.DateTimeEvent,
		&qgp.Latitude,
		&qgp.Longitude,
		&qgp.Speed,
		&qgp.Direction,
		&qgp.GpsState,
		&qgp.SecondsSinceLastPosition,
		&qgp.StateDigitalInputs,
		&qgp.EventNumber,
		&qgp.DilutionHorizontalPrecision)
	if err != nil {
		return &QGP{}
	}

	return qgp
}

func NewQGP(ms Message) VirlocReport {
	qgp := &QGP{
		RawData: ms.Message,
		Header:  ms.Header,
	}
	return qgp
}
