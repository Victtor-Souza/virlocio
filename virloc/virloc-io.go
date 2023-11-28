package virloc

import (
	"fmt"
	"strings"
)

func Serialize(message string, v VirlocReport) error {
	v, err := v.serialize(message)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func NewVirlocReport(ms string) VirlocReport {
	hd := extractHeader(ms)
	message := message{
		RawData: ms,
		Header:  hd,
	}
	switch getReportType(ms) {
	case REPORT_RGP:
		return newQGP(message)
	case REPORT_RTT:
		return newQTT(message)
	case REPORT_RSD:
		return newQSD(message)
	case REPORT_RUV00:
		return newRUV00(message)
	case REPORT_RUV01:
		return newRUV01(message)
	case REPORT_RUV03:
		return newRUV03(message)
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

	if reportType == "RUV" {
		_, err := fmt.Sscanf(removeSpecialCharsAndSpaces(ms), "%5s", &reportType)
		if err != nil {
			return ""
		}

		return ReportType(reportType)
	}

	return ReportType(reportType)
}
