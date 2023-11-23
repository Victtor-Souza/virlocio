package virloc

type Message struct {
	Message string
	Header  Header
}

type Header struct {
	DeviceId      string
	MessageNumber string
	CheckSum      string
}

type VirlocReport interface {
	ToRawMessage() string
	AcceptMessage() string
	serialize(msg string) VirlocReport
}
