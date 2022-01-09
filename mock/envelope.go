package mock

import (
	"github.com/riid/messenger"
	"github.com/stretchr/testify/mock"
)

type Envelope struct {
	mock.Mock
}

func (m *Envelope) Message() interface{} {
	return m.Called().Get(0)
}

func (m *Envelope) Headers() map[string][]string {
	return m.Called().Get(0).(map[string][]string)
}

func (m *Envelope) Header(name string) []string {
	return m.Called(name).Get(0).([]string)
}

func (m *Envelope) HasHeader(name string) bool {
	return m.Called(name).Bool(0)
}

func (m *Envelope) LastHeader(name string) (string, bool) {
	args := m.Called(name)
	return args.String(0), args.Bool(1)
}

func (m *Envelope) FirstHeader(name string) (string, bool) {
	args := m.Called(name)
	return args.String(0), args.Bool(1)
}

func (m *Envelope) Is(e messenger.Envelope) bool {
	return m.Called(e).Bool(0)
}
