package filedirs

import (
	"errors"
	"io"
	"os"
)

// Operate manipulates files and directories
func Operate() error {
	err := os.Mkdir("example_dir", os.FileMode(0755))
	if err != nil {
		return err
	}
	// go to the /tmp directory
	if err := os.Chdir("example_dir"); err != nil {
		return err
	}
	// f is a generic file object it also implements multiple interfaces and can be used as a reader or writer
	f, err := os.Create("text.txt")
	if err != nil {
		return err
	}

	// we write a known-length value to the file and validate that it wrote correctly
	value := []byte("hello")
	count, err := f.Write(value)
	if err != nil {
		return err
	}
	if count != len(value) {
		return errors.New("incorrect length returned from write")
	}
	if err := f.Close(); err != nil {
		return err
	}
	// read the file
	f, err = os.Open("text.txt")
	if err != nil {
		return err
	}
	io.Copy(os.Stdout, f)
	if err := f.Close(); err != nil {
		return err
	}
	// go the /tmp directory
	if err := os.Chdir(".."); err != nil {
		return err
	}
	// clianup, os.RemoveAll can be dangerous if you point at the wrong directory, use
	// user input, and especially if you run as root
	if err := os.RemoveAll("example_dir"); err != nil {
		return err
	}
	return nil
}
