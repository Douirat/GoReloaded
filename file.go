package Reload

import "os"

// Declare a file object:
type File struct {
	Name, Content string
}

// Instantiate the file object:
func NewFile(name, content string) *File {
	new_file := new(File)
	new_file.Name = name
	new_file.Content = content
	return new_file
}

// Read data from an existing new file:
func (file *File) Read() {
	data, err := os.ReadFile(file.Name)
	if err != nil {
		panic(err)
	}
	file.Content = string(data)
}

// Create a result new file or overwrite an existing with a result extracted from the sample file:
func (file *File) CopyData() {
	f, err := os.Create(file.Name)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, Err := f.WriteString(file.Content)
	if Err != nil {
		panic(Err)
	}
}
