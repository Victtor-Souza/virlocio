package virloc

import (
	"fmt"
)

type QTT struct {
	RawData                            string
	PackageType                        string
	Date                               string
	Time                               string
	Latitude                           string
	Longitude                          string
	Speed                              string
	DirectionLastPosition              string
	LastPositionState                  string
	SecondsSinceLastPosition           string
	StateDigitalInputs                 string
	EventNumber                        string
	DilutionHorizontalPrecision        string
	ExternalGPSStatus                  string
	PositionalGPSStatus                string
	NumberOfSatellitesToFixGPSPosition string
	NumberOfSatellitesUsedByGPSNow     string
	NumberOfSatellitesAvailable        string
	Output0State                       string
	Output1State                       string
	AnalogInput0Voltage                string
	AnalogInput1Voltage                string
	AnalogInput2Voltage                string
	InternalBatteryVoltage             string
	MainSupplyVoltage                  string
	Header                             Header
}

func NewQTT(ms Message) VirlocReport {
	qtt := QTT{
		RawData: ms.Message,
		Header:  ms.Header,
	}
	return &qtt
}

func (qtt *QTT) ToRawMessage() string {
	return qtt.RawData
}

func (qtt *QTT) AcceptMessage() string {
	ackmsg := fmt.Sprintf(">ACK;%s;%s", qtt.Header.DeviceId, qtt.Header.MessageNumber)
	chksum := CalculateChecksum(qtt.RawData)
	return fmt.Sprintf("%s;*%s<", ackmsg, chksum)
}

func (qtt *QTT) serialize(msg string) VirlocReport {
	messagewspace := RemoveSpecialCharsAndSpaces(msg)
	if _, err := fmt.Sscanf(messagewspace, "%3s%6s%6s%8s%9s%3s%3s%1s%2s%2s%2s%2s%1s%1s%2s%2s%2s%1s%1s%1s%4s%4s%4s%4s%4s",
		&qtt.PackageType,
		&qtt.Date,
		&qtt.Time,
		&qtt.Latitude,
		&qtt.Longitude,
		&qtt.Speed,
		&qtt.DirectionLastPosition,
		&qtt.LastPositionState,
		&qtt.PositionalGPSStatus,
		&qtt.SecondsSinceLastPosition,
		&qtt.StateDigitalInputs,
		&qtt.EventNumber,
		&qtt.DilutionHorizontalPrecision,
		&qtt.ExternalGPSStatus,
		&qtt.PositionalGPSStatus,
		&qtt.NumberOfSatellitesToFixGPSPosition,
		&qtt.NumberOfSatellitesUsedByGPSNow,
		&qtt.NumberOfSatellitesAvailable,
		&qtt.Output0State,
		&qtt.Output1State,
		&qtt.AnalogInput0Voltage,
		&qtt.AnalogInput1Voltage,
		&qtt.AnalogInput2Voltage,
		&qtt.InternalBatteryVoltage,
		&qtt.MainSupplyVoltage,
	); err != nil {
		return &QTT{}
	}

	return qtt
}
