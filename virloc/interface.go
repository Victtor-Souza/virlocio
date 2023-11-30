package virloc

import (
	"errors"
	"fmt"
	"strings"
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

func newHeader(deviceId, messageNumber, checkSum string) header {
	return header{
		DeviceId:      strings.Replace(deviceId, "ID=", "", 1),
		MessageNumber: messageNumber,
		CheckSum:      checkSum,
	}
}

func newMessage(msg string) message {
	messageArr := strings.Split(msg, ";")
	return message{
		RawData: messageArr[0],
		Header:  newHeader(messageArr[1], messageArr[2], messageArr[3]),
	}
}

type VirlocReport interface {
	GetSerialNumber() string
	ToRawMessage() string
	AcceptMessage() string
	serialize(msg string) (VirlocReport, error)
}

func (m *message) AcceptMessage() string {
	ackmsg := fmt.Sprintf(">ACK;%s;%s", m.Header.DeviceId, m.Header.MessageNumber)
	chksum := CalculateChecksum(m.RawData)
	return fmt.Sprintf("%s;*%s<", ackmsg, chksum)
}

func (m *message) GetSerialNumber() string {
	return m.Header.DeviceId
}
