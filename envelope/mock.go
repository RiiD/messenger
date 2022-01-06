package envelope

import (
	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) Message() interface{} {
	return m.Called().Get(0)
}

func (m *Mock) Headers() map[string][]string {
	return m.Called().Get(0).(map[string][]string)
}

func (m *Mock) Header(name string) []string {
	return m.Called(name).Get(0).([]string)
}

func (m *Mock) HasHeader(name string) bool {
	return m.Called(name).Bool(0)
}

func (m *Mock) LastHeader(name string) (string, bool) {
	args := m.Called(name)
	return args.String(0), args.Bool(1)
}

func (m *Mock) FirstHeader(name string) (string, bool) {
	args := m.Called(name)
	return args.String(0), args.Bool(1)
}

func (m *Mock) Is(e Envelope) bool {
	return m.Called(e).Bool(0)
}
