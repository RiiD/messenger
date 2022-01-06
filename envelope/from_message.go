package envelope

func FromMessage(message interface{}) *fromMessage {
	return &fromMessage{
		message: message,
	}
}

type fromMessage struct {
	message interface{}
}

func (e *fromMessage) Message() interface{} {
	return e.message
}

func (e *fromMessage) Headers() map[string][]string {
	return map[string][]string{}
}

func (e *fromMessage) Header(_ string) []string {
	return []string{}
}

func (e *fromMessage) HasHeader(_ string) bool {
	return false
}

func (e *fromMessage) LastHeader(_ string) (string, bool) {
	return "", false
}

func (e *fromMessage) FirstHeader(_ string) (string, bool) {
	return "", false
}

func (e *fromMessage) Is(other Envelope) bool {
	return e == other
}
