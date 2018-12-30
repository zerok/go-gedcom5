//go:generate goyacc -p "line" -o "xx_line.go" gedcom5.y
package gedcom5

type Header struct{}

type Trailer struct{}

type Record interface{}

type File struct {
	Lines   []Line
	Header  Header
	Records []Record
	Trailer Trailer
}
