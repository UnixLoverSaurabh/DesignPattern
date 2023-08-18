package directory

import "fmt"

type File struct {
	Name string
}

func (f *File) Search() {
	fmt.Println("This is a file", f.getFileName())
}

func NewFile(name string) *File {
	return &File{Name: name}
}

func (f *File) getFileName() string {
	return f.Name
}
