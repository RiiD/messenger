package matcher

import (
	"github.com/riid/messenger"
	"reflect"
)

// Type matches envelopes with message with the same type as the prototype
func Type(prototype interface{}) messenger.Matcher {
	t := reflect.TypeOf(prototype)
	return MatchFunc(func(e messenger.Envelope) bool {
		return t == reflect.TypeOf(e.Message())
	})
}
