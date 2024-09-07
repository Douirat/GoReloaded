package reload

import (
	"io"
	"os"
)

// Declare a new file object
type File struct {
	Name string
	Content string
}

// Instantiate a new object:
type Filer interface {
	WriteInto()
	CopyData() string
}

// A function to instantiate a new file in go:
func NewFile(fn, cont string) *File {
	return &File{Name: fn, Content: cont}
}

// A Method to create and insert data into a file:
func (f *File) WriteInto() {
	// Create a file using the OS package:
	file, err := os.Create(f.Name)
	HandleNilError(err)
	_, Err := io.WriteString(file, f.Content)
	HandleNilError(Err)
	defer file.Close()
}

// Write a Method to read data from a file:
func (f *File)CopyData() string {
	data, err := os.ReadFile(f.Name)
	HandleNilError(err)
	return string(data)
}

// A funnction to handle nil errors:
func HandleNilError(err error) {
	if err != nil {
		panic(err)
	}
}

