package virloc

import (
	"fmt"
	"strings"
)

func Serialize(message string, v VirlocReport) error {
	v = v.serialize(message)
	return nil
}

func NewVirlocReport(ms string) VirlocReport {
	hd := extractHeader(ms)
	message := message{
		Message: ms,
		Header:  hd,
	}
	switch getReportType(ms) {
	case REPORT_RGP:
		return newQGP(message)
	case REPORT_RTT:
		return newQTT(message)
	case REPORT_RSD:
		return newQSD(message)
	default:
		return nil
	}
}

func extractHeader(message string) header {
	messageArr := strings.Split(message, ";")
	return header{
		DeviceId:      messageArr[1],
		MessageNumber: messageArr[2],
		CheckSum:      messageArr[3],
	}
}

func getReportType(ms string) ReportType {
	var reportType string

	_, err := fmt.Sscanf(removeSpecialCharsAndSpaces(ms), "%3s", &reportType)
	if err != nil {
		return ""
	}

	return ReportType(reportType)
}
