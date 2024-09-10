package ofctl

import (
	"bytes"
	"io"
	"os"
	"syscall"
)

//#include <stdio.h>
import "C"

func newStub(std int) (stub *stubStd, err error) {
	stub = &stubStd{std: std}

	err = stub.init()
	if err != nil {
		return nil, err
	}

	err = stub.spoof()
	if err != nil {
		return nil, err
	}

	return stub, nil
}

// stubStd inspired by:
// https://github.com/eliben/code-for-blog/blob/main/2020/go-fake-stdio/snippets/redirect-cgo-stdout.go
type stubStd struct {
	std int // syscall.Stdout || syscall.Stderr
	fd  int // file descriptor of new stubStd

	// read and write connected pairs of os.Pipe
	read  *os.File
	write *os.File

	buffer chan []byte
}

func (stub *stubStd) init() (err error) {
	// Clone std to fd
	stub.fd, err = syscall.Dup(stub.std)
	if err != nil {
		return err
	}

	// Open pipe for stubStd
	stub.read, stub.write, err = os.Pipe()
	if err != nil {
		return err
	}

	return nil
}

func (stub *stubStd) spoof() (err error) {
	// Clone the pipe's writer to the actual Stdout descriptor; from this point
	// on, writes to Stdout will go to w.
	if err = syscall.Dup2(int(stub.write.Fd()), stub.std); err != nil {
		return err
	}

	// Background goroutine that drains the reading end of the pipe.
	stub.buffer = make(chan []byte)
	go func() {
		var b bytes.Buffer
		io.Copy(&b, stub.read)
		stub.buffer <- b.Bytes()
	}()

	return nil
}

func (stub *stubStd) close() {
	stub.write.Close()
	syscall.Close(stub.std)

	// Restore original Stdout.
	syscall.Dup2(stub.fd, stub.std)
	syscall.Close(stub.fd)
}

func (stub *stubStd) String() string {
	stub.close()

	// Rendezvous with the reading goroutine.
	return string(<-stub.buffer)
}

func catchStd(f func()) (stdout, stderr string) {
	stubStdout, err := newStub(syscall.Stdout)
	if err != nil {
		panic(err)
	}

	stubStderr, err := newStub(syscall.Stderr)
	if err != nil {
		panic(err)
	}

	// ----> The actual cgo call <----
	f()

	// Cleanup
	C.fflush(nil)

	return stubStdout.String(), stubStderr.String()
}
