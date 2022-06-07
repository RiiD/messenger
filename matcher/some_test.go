package matcher

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/riid/messenger"
	"github.com/riid/messenger/envelope"
)

func TestSome(t *testing.T) {
	positive := &matcherStub{true}
	negative := &matcherStub{false}
	e := envelope.FromMessage([]byte{})

	type testCase struct {
		matchers []messenger.Matcher
		expected bool
	}

	suite := map[string]testCase{
		"no matchers": {
			expected: false,
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
			expected: true,
		},
	}

	for name, tc := range suite {
		tc := tc
		t.Run(
			name, func(t *testing.T) {
				t.Parallel()

				m := Some(tc.matchers...)
				assert.Equal(t, tc.expected, m.Matches(e))
			},
		)
	}
}
