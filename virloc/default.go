package virloc

type (
	Default struct {
		message
	}
)

// ToRawMessage implements VirlocReport.
func (r *Default) ToRawMessage() string {
	return r.RawData
}

// serialize implements VirlocReport.
func (r *Default) serialize(msg string) (VirlocReport, error) {
	return r, nil
}

func newDefault(ms message) VirlocReport {
	return &Default{
		message: ms,
	}
}
