//go:generate python3 generators/file_record_lookup.py --outfile file_gen_lookup.go --package gedcom5 --recordtypes Individual Family
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
