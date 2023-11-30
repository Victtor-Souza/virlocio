package virloc

type ReportType string
type CommandName string
type CommandMessage string

const (
	REPORT_RTT   ReportType = "RTT"
	REPORT_RGP   ReportType = "RGP"
	REPORT_RSD   ReportType = "RSD"
	REPORT_RUV01 ReportType = "RUV01"
	REPORT_RUV00 ReportType = "RUV00"
	REPORT_RUV02 ReportType = "RUV02"
	REPORT_RUV03 ReportType = "RUV03"

	TURN_OFF_OUT0 CommandName = "TURN_OFF_OUT0"
	TURN_ON_OUT0  CommandName = "TURN_ON_OUT0"
	TURN_ON_OUT1  CommandName = "TURN_ON_OUT1"
	TURN_OFF_OUT1 CommandName = "TURN_OFF_OUT1"

	TURN_OFF_OUT0MSG CommandMessage = "SSXP00"
	TURN_ON_OUT0MSG  CommandMessage = "SSXP01"

	TURN_ON_OUT1MSG  CommandMessage = "SSXP11"
	TURN_OFF_OUT1MSG CommandMessage = "SSXP10"
)
