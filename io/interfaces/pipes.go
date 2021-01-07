package interfaces

import (
	"io"
	"os"
)

// PipeExample helps give some more examples of using io interfaces
func PipeExample() error {
	// the pipe reader and pipe writer implement
	r, w := io.Pipe()
	go func() {
		w.Write([]byte("test\n"))
		w.Close()
	}()
	if _, err := io.Copy(os.Stdout, r); err != nil {
		return err
	}
	return nil
}
