package envelope

func WithMessage(e Envelope, body interface{}) *withMessage {
	return &withMessage{
		message: body,
		wrapped: e,
	}
}

type withMessage struct {
	message interface{}
	wrapped Envelope
}

func (w *withMessage) Message() interface{} {
	return w.message
}

func (w *withMessage) Headers() map[string][]string {
	return w.wrapped.Headers()
}

func (w *withMessage) Header(name string) []string {
	return w.wrapped.Header(name)
}

func (w *withMessage) HasHeader(name string) bool {
	return w.wrapped.HasHeader(name)
}

func (w *withMessage) LastHeader(name string) (string, bool) {
	return w.wrapped.LastHeader(name)
}

func (w *withMessage) FirstHeader(name string) (string, bool) {
	return w.wrapped.FirstHeader(name)
}

func (w *withMessage) Is(e Envelope) bool {
	if e == w {
		return true
	}
	return w.wrapped.Is(e)
}
