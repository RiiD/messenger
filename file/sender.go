package file

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/riid/messenger/envelope"
	"os"
)

func Sender(path string, append bool) (*sender, error) {
	flags := os.O_CREATE | os.O_WRONLY
	if append {
		flags = flags | os.O_APPEND
	} else {
		flags = flags | os.O_TRUNC
	}

	file, err := os.OpenFile(path, flags, 0666)
	if err != nil {
		return nil, err
	}

	return &sender{file: file}, nil
}

type sender struct {
	file *os.File
}

func (s *sender) Send(_ context.Context, e envelope.Envelope) error {
	l := &line{
		Headers: e.Headers(),
		Body:    e.Message().([]byte),
	}

	marshalledLine, err := json.Marshal(l)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(s.file, string(marshalledLine))
	if err != nil {
		return err
	}

	return nil
}
