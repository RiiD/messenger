package messenger

type Matcher interface {
	Matches(e Envelope) bool
}
