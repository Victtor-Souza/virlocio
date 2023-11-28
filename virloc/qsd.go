package virloc

import "fmt"

type QSD struct {
	message
	PackageType string
	Speed       string
}

func (qsd *QSD) ToRawMessage() string {
	return qsd.RawData
}

func (qsd *QSD) serialize(msg string) (VirlocReport, error) {
	msgw := removeSpecialCharsAndSpaces(msg)

	_, err := fmt.Sscanf(msgw, "%3s%4s", &qsd.PackageType, &qsd.Speed)
	if err != nil {
		return nil, ErrReadingMessage(err)
	}

	return qsd, nil
}

func newQSD(ms message) VirlocReport {
	qsd := &QSD{
		message: ms,
	}

	return qsd
}
