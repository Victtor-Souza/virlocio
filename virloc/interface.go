package virloc

type message struct {
	Message string
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
	serialize(msg string) VirlocReport
}
