package virloc

import (
	"errors"
	"fmt"
)

var (
	ErrReportNotConfigured = func(s string) error { return errors.New(fmt.Sprintf("REPORT NOT CONFIGURED: %s", s)) }
)

func Serialize(message string, v VirlocReport) error {
	v, err := v.serialize(message)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func NewVirlocReport(ms string) (VirlocReport, error) {
	message := newMessage(ms)
	rp := getReportType(ms)
	switch rp {
	case REPORT_RGP:
		return newQGP(message), nil
	case REPORT_RTT:
		return newQTT(message), nil
	case REPORT_RSD:
		return newQSD(message), nil
	case REPORT_RUV00:
		return newRUV00(message), nil
	case REPORT_RUV01:
		return newRUV01(message), nil
	case REPORT_RUV02:
		return newRUV02(message), nil
	case REPORT_RUV03:
		return newRUV03(message), nil
	default:
		return nil, ErrReportNotConfigured(string(rp))
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
