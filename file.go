package gedcom5

type Header struct {
	Lines []Line
}

type Trailer struct {
	Lines []Line
}

type File struct {
	Lines   []Line
	Header  Header
	Records []Record
	Trailer Trailer
}

func NewFile() *File {
	f := File{
		Lines:   make([]Line, 0, 10),
		Records: make([]Record, 0, 10),
	}
	return &f
}
