package envelope

type Envelope interface {
	Message() interface{}
	Headers() map[string][]string
	Header(name string) []string
	HasHeader(name string) bool
	LastHeader(name string) (string, bool)
	FirstHeader(name string) (string, bool)
	Is(e Envelope) bool
}
