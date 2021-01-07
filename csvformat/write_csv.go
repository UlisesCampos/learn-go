package csvformat

import (
	"bytes"
	"encoding/csv"
	"io"
	"os"
)

// A Book has an Author and Title
type Book struct {
	Author string
	Title  string
}

// Books is a named type for an array of books
type Books []Book

// ToCSV takes a set of Books and writes to an io.Writer it returns any errors
func (books *Books) ToCSV(w io.Writer) error {
	n := csv.NewWriter(w)
	err := n.Write([]string{"Author", "Title"})
	if err != nil {
		return err
	}
	for _, book := range *books {
		err := n.Write([]string{book.Author, book.Title})
		if err != nil {
			return err
		}
	}
	n.Flush()
	return n.Error()
}

// WriteCSVOutput returns a buffer csv for a set of books
func WriteCSVOutput() error {
	b := Books{
		Book{
			Author: "George Orwell",
			Title:  "1984",
		},
		Book{
			Author: "Aldous Huxley",
			Title:  "Un mundo feliz",
		},
	}
	return b.ToCSV(os.Stdout)
}

// WiteCSVBuffer return a buffer csv for a set of books
func WiteCSVBuffer() (*bytes.Buffer, error) {
	b := Books{
		Book{
			Author: "Miguel de Cervantes",
			Title:  "Don Quijote de la Mancha",
		},
		Book{
			Author: "Marqués de Sade",
			Title:  "Justina",
		},
	}
	w := &bytes.Buffer{}
	err := b.ToCSV(w)
	return w, err
}
