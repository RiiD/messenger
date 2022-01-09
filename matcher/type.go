package matcher

import (
	"github.com/riid/messenger"
	"reflect"
)

func Type(v interface{}) messenger.Matcher {
	t := reflect.TypeOf(v)
	return MatchFunc(func(e messenger.Envelope) bool {
		return t == reflect.TypeOf(e.Message())
	})
}
