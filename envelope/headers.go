package envelope

func WithHeaders(wrapped Envelope, headers map[string][]string) *withHeaders {
	return &withHeaders{
		wrapped: wrapped,
		headers: headers,
	}
}

type withHeaders struct {
	wrapped Envelope
	headers map[string][]string
}

func (w *withHeaders) Message() interface{} {
	return w.wrapped.Message()
}

func (w *withHeaders) Headers() map[string][]string {
	wrappedHeaders := w.wrapped.Headers()
	for name, hh := range w.headers {
		headers, found := wrappedHeaders[name]
		if found {
			headers = append(headers, hh...)
		} else {
			headers = hh
		}
		wrappedHeaders[name] = headers
	}

	return wrappedHeaders
}

func (w *withHeaders) Header(name string) []string {
	wrappedHeaders := w.wrapped.Header(name)
	headers, found := w.headers[name]
	if found {
		return append(wrappedHeaders, headers...)
	}
	return wrappedHeaders
}

func (w *withHeaders) HasHeader(name string) bool {
	_, found := w.headers[name]
	return found || w.wrapped.HasHeader(name)
}

func (w *withHeaders) LastHeader(name string) (string, bool) {
	values, found := w.headers[name]
	if found && len(values) > 0 {
		return values[len(values)-1], true
	}
	return w.wrapped.LastHeader(name)
}

func (w *withHeaders) FirstHeader(name string) (string, bool) {
	value, found := w.wrapped.FirstHeader(name)
	if found {
		return value, found
	}
	values, found := w.headers[name]
	if !found || len(values) == 0 {
		return "", false
	}

	return values[0], true
}

func (w *withHeaders) Is(e Envelope) bool {
	return w == e || w.wrapped.Is(e)
}
