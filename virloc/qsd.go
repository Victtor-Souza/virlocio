package virloc

import "fmt"

type QSD struct {
	RawData     string
	Header      header
	PackageType string
	Speed       string
}

func (qsd *QSD) AcceptMessage() string {
	ackmsg := fmt.Sprintf(">ACK;%s;%s", qsd.Header.DeviceId, qsd.Header.MessageNumber)
	chksum := calculateChecksum(qsd.RawData)
	return fmt.Sprintf("%s;*%s<", ackmsg, chksum)
}

func (qsd *QSD) ToRawMessage() string {
	return qsd.RawData
}

func (qsd *QSD) serialize(msg string) VirlocReport {
	msgw := removeSpecialCharsAndSpaces(msg)

	_, err := fmt.Sscanf(msgw, "%3s%4s", &qsd.PackageType, &qsd.Speed)
	if err != nil {
		return &QSD{}
	}

	return qsd
}

func newQSD(ms message) VirlocReport {
	qsd := &QSD{
		RawData: ms.Message,
		Header:  ms.Header,
	}

	return qsd
}
