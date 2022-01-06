package file

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"syscall"
)

func Follower() *follower {
	return &follower{}
}

type follower struct {
}

func (f follower) Follow(ctx context.Context, filename string) (<-chan []byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	fd, err := syscall.InotifyInit()
	if err != nil {
		return nil, err
	}
	wd, err := syscall.InotifyAddWatch(fd, filename, syscall.IN_MODIFY)
	if err != nil {
		return nil, err
	}

	go func() {
		<-ctx.Done()
		_, _ = syscall.InotifyRmWatch(fd, uint32(wd))
		_ = file.Close()
	}()

	ch := make(chan []byte, 0)

	go func() {
		for {
			var by []byte
			_, err := fmt.Fscanln(file, &by)
			if err != nil && err != io.EOF {
				close(ch)
				break
			}
			if err == io.EOF {
				if err = waitForChange(fd); err != nil {
					close(ch)
					break
				}
				continue
			}

			ch <- by
		}
	}()

	return ch, nil
}

func waitForChange(fd int) error {
	for {
		var buf [syscall.SizeofInotifyEvent]byte
		_, err := syscall.Read(fd, buf[:])
		if err != nil {
			return err
		}
		r := bytes.NewReader(buf[:])
		var ev = syscall.InotifyEvent{}
		_ = binary.Read(r, binary.LittleEndian, &ev)
		if ev.Mask&syscall.IN_MODIFY == syscall.IN_MODIFY {
			return nil
		}
	}
}
