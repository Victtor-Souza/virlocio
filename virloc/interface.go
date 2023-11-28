package virloc

import (
	"errors"
	"fmt"
)

var (
	ErrReadingMessage = func(err error) error {
		return errors.New("ERROR READING MESSAGE: " + err.Error())
	}
)

type message struct {
	RawData string
	Header  header
}

type header struct {
	DeviceId      string
	MessageNumber string
	CheckSum      string
}

type VirlocReport interface {
	ToRawMessage() string
	AcceptMessage() string
	serialize(msg string) (VirlocReport, error)
}

func (m *message) AcceptMessage() string {
	ackmsg := fmt.Sprintf(">ACK;%s;%s", m.Header.DeviceId, m.Header.MessageNumber)
	chksum := calculateChecksum(m.RawData)
	return fmt.Sprintf("%s;*%s<", ackmsg, chksum)
}
