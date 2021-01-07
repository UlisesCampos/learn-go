package filedirs

import (
	"bytes"
	"io"
	"os"
	"strings"
)

// Capitalizer opens a file, reads the contents, then writes those contents to a second file
func Capitalizer(f1 *os.File, f2 *os.File) error {
	if _, err := f1.Seek(0, io.SeekStart); err != nil {
		return err
	}
	var tmp = new(bytes.Buffer)
	if _, err := io.Copy(tmp, f1); err != nil {
		return err
	}
	s := strings.ToUpper(tmp.String())
	if _, err := io.Copy(f2, strings.NewReader(s)); err != nil {
		return err
	}
	return nil
}
