package virloc

import (
	"fmt"
	"strings"
)

type RUV00 struct {
	message
	PackageType                 string
	EventIndexDispatch          string
	ProtocolIdentifier          string
	Date                        string
	Time                        string
	BackupBatteryVoltage        string
	MainSupplyVoltage           string
	VIRLOCSerialNumber          string
	VIRLOCFirmware              string
	VIRLOCPlate                 string
	ScriptVersion               string
	SerialPeripheral            string
	SpeedEventInRain            string
	HourmeterSelection          string
	AccelerationEventsSelection string
	ICCIDChip                   string
}

// ToRawMessage implements VirlocReport.
func (r *RUV00) ToRawMessage() string {
	return r.RawData
}

// serialize implements VirlocReport.
func (r *RUV00) serialize(msg string) (VirlocReport, error) {
	messagewsp := removeSpecialCharsAndSpaces(msg)
	msgarr := strings.Split(messagewsp, ",")

	if _, err := fmt.Sscanf(msgarr[0], "%5s%3s", &r.PackageType, &r.EventIndexDispatch); err != nil {
		return nil, ErrReadingMessage(err)
	}

	r.ProtocolIdentifier = msgarr[1]

	if _, err := fmt.Sscanf(msgarr[2], "%6s%6s", &r.Date, &r.Time); err != nil {
		return nil, ErrReadingMessage(err)
	}

	if _, err := fmt.Sscanf(msgarr[3], "%4s%4s", &r.BackupBatteryVoltage, &r.MainSupplyVoltage); err != nil {
		return nil, ErrReadingMessage(err)
	}

	r.VIRLOCSerialNumber = msgarr[4]
	r.VIRLOCFirmware = msgarr[5]
	r.VIRLOCPlate = msgarr[6]
	r.ScriptVersion = msgarr[7]
	r.SerialPeripheral = msgarr[8]
	r.SpeedEventInRain = msgarr[9]
	r.HourmeterSelection = msgarr[10]
	r.ICCIDChip = removeDeviceData(msgarr[11])

	return r, nil
}

func newRUV00(ms message) VirlocReport {
	return &RUV00{
		message: ms,
	}
}
