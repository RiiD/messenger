package matcher

import (
	"github.com/riid/messenger/envelope"
	"reflect"
)

func Type(v interface{}) Matcher {
	t := reflect.TypeOf(v)
	return MatchFunc(func(e envelope.Envelope) bool {
		return t == reflect.TypeOf(e.Message())
	})
}
