package matcher

import (
	"github.com/riid/messenger"
	"github.com/riid/messenger/envelope"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestType_Matches(t *testing.T) {
	tcc := map[string]struct {
		t   interface{}
		e   messenger.Envelope
		res bool
	}{
		"when type is string and message is a string should return true": {
			t:   "",
			e:   envelope.FromMessage("test"),
			res: true,
		},
		"when type is int and message is an int should return true": {
			t:   1,
			e:   envelope.FromMessage(99),
			res: true,
		},
		"when type is []byte and message is []byte should return true": {
			t:   []byte{},
			e:   envelope.FromMessage([]byte("test")),
			res: true,
		},
		"when type is string and message is an int should return false": {
			t:   "",
			e:   envelope.FromMessage(123),
			res: false,
		},
	}

	for name, tc := range tcc {
		t.Run(name, func(t *testing.T) {
			m := Type(tc.t)

			res := m.Matches(tc.e)

			assert.Equal(t, tc.res, res)
		})
	}
}
