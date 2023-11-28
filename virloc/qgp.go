package virloc

import (
	"fmt"
)

type QGP struct {
	message
	PackageType                 string
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

func (qgp *QGP) ToRawMessage() string {
	return qgp.RawData
}

func (qgp *QGP) serialize(msg string) (VirlocReport, error) {
	messagewspace := removeSpecialCharsAndSpaces(msg)
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
		return nil, ErrReadingMessage(err)
	}

	return qgp, nil
}

func newQGP(ms message) VirlocReport {
	qgp := &QGP{
		message: ms,
	}
	return qgp
}
