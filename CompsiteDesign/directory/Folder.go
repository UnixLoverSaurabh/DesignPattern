package directory

type Folder struct {
	name        string
	Directories []Directory
}

func NewFolder(name string) *Folder {
	return &Folder{name: name}
}

func (f *Folder) Search() {
	for _, folder := range f.Directories {
		folder.Search()
	}
}

func (f *Folder) Add(fileOrFolder Directory) {
	f.Directories = append(f.Directories, fileOrFolder)
}
