package main

import (
	"DesignPattern/CompsiteDesign/directory"
	"fmt"
)

/*
Folder2

	File2
	File3
	Folder1
		File1
*/

func main() {
	file1 := directory.NewFile("File1")
	file2 := directory.NewFile("File2")
	file3 := directory.NewFile("File3")

	folder1 := directory.NewFolder("folder1")
	folder2 := directory.NewFolder("folder2")

	folder1.Add(file1)
	folder2.Add(file2)
	folder2.Add(file3)
	folder2.Add(folder1)

	folder1.Search()
	fmt.Println("--------------------------")
	folder2.Search()
}
