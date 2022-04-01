package matcher

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/riid/messenger"
	"github.com/riid/messenger/envelope"
)

type matcherStub struct {
	ret bool
}

func (m *matcherStub) Matches(_ messenger.Envelope) bool {
	return m.ret
}

func TestAll(t *testing.T) {
	positive := &matcherStub{true}
	negative := &matcherStub{false}
	e := envelope.FromMessage([]byte{})

	type testCase struct {
		matchers []messenger.Matcher
		expected bool
	}

	suite := map[string]testCase{
		"no matchers": {
			expected: true,
		},
		"all positive matchers": {
			matchers: []messenger.Matcher{
				positive,
				positive,
			},
			expected: true,
		},
		"all negative matchers": {
			matchers: []messenger.Matcher{
				negative,
				negative,
			},
			expected: false,
		},
		"one negative matcher": {
			matchers: []messenger.Matcher{
				positive,
				negative,
				positive,
			},
			expected: false,
		},
	}

	for name, tc := range suite {
		tc := tc
		t.Run(
			name, func(t *testing.T) {
				t.Parallel()

				m := All(tc.matchers...)
				assert.Equal(t, tc.expected, m.Matches(e))
			},
		)
	}
}
