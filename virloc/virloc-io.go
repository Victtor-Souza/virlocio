package virloc

import (
	"fmt"
	"strings"
)

func Serialize(message string, v VirlocReport) error {
	v = v.serialize(message)
	return nil
}

func extractHeader(message string) Header {
	messageArr := strings.Split(message, ";")
	return Header{
		DeviceId:      messageArr[1],
		MessageNumber: messageArr[2],
		CheckSum:      messageArr[3],
	}
}

func NewVirlocReport(ms string) VirlocReport {
	hd := extractHeader(ms)
	message := Message{
		Message: ms,
		Header:  hd,
	}
	switch getReportType(ms) {
	case REPORT_RGP:
		return NewQGP(message)
	case REPORT_RTT:
		return NewQTT(message)
	case REPORT_RSD:
		return NewQSD(message)
	default:
		return nil
	}
}

func getReportType(ms string) ReportType {
	var reportType string

	_, err := fmt.Sscanf(RemoveSpecialCharsAndSpaces(ms), "%3s", &reportType)
	if err != nil {
		return ""
	}

	return ReportType(reportType)
}
