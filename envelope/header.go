package envelope

import "github.com/riid/messenger"

func WithHeader(wrapped messenger.Envelope, key, value string) *withHeader {
	return &withHeader{
		wrapped: wrapped,
		name:    key,
		value:   value,
	}
}

type withHeader struct {
	wrapped messenger.Envelope
	name    string
	value   string
}

func (e *withHeader) Message() interface{} {
	return e.wrapped.Message()
}

func (e *withHeader) Headers() map[string][]string {
	allHeaders := e.wrapped.Headers()
	headers, ok := allHeaders[e.name]
	if ok {
		headers = append(headers, e.value)
	} else {
		headers = []string{e.value}
	}

	allHeaders[e.name] = headers
	return allHeaders
}

func (e *withHeader) Header(key string) []string {
	headers := e.wrapped.Header(key)
	if key == e.name {
		headers = append(headers, e.value)
	}
	return headers
}

func (e *withHeader) HasHeader(name string) bool {
	return name == e.name || e.wrapped.HasHeader(name)
}

func (e *withHeader) LastHeader(name string) (string, bool) {
	if name == e.name {
		return e.value, true
	}

	return e.wrapped.LastHeader(name)
}

func (e *withHeader) FirstHeader(name string) (string, bool) {
	if v, f := e.wrapped.FirstHeader(name); f {
		return v, f
	}
	if e.name != name {
		return "", false
	}

	return e.value, true
}

func (e *withHeader) Is(other messenger.Envelope) bool {
	if e == other {
		return true
	}

	return e.wrapped.Is(other)
}

func WithoutHeader(e messenger.Envelope, name string) *withoutHeader {
	return &withoutHeader{
		wrapped: e,
		name:    name,
	}
}

type withoutHeader struct {
	wrapped messenger.Envelope
	name    string
}

func (w *withoutHeader) Message() interface{} {
	return w.wrapped.Message()
}

func (w *withoutHeader) Headers() map[string][]string {
	headers := w.wrapped.Headers()
	delete(headers, w.name)
	return headers
}

func (w *withoutHeader) Header(name string) []string {
	if name != w.name {
		return w.wrapped.Header(name)
	}

	return []string{}
}

func (w *withoutHeader) HasHeader(name string) bool {
	if name != w.name {
		return w.wrapped.HasHeader(name)
	}

	return false
}

func (w *withoutHeader) LastHeader(name string) (string, bool) {
	if name != w.name {
		return w.wrapped.LastHeader(name)
	}

	return "", false
}

func (w *withoutHeader) FirstHeader(name string) (string, bool) {
	if name != w.name {
		return w.wrapped.FirstHeader(name)
	}

	return "", false
}

func (w *withoutHeader) Is(other messenger.Envelope) bool {
	if w == other {
		return true
	}

	return w.wrapped.Is(other)
}
