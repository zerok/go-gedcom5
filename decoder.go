package gedcom5

import "io"

type Decoder struct {
	p *FileParser
}

func NewDecoder(in io.Reader) *Decoder {
	s := NewScanner(in)
	return &Decoder{p: NewFileParser(s)}
}

func (ld *Decoder) Decode(out *File) error {
	f, err := ld.p.ParseFile()
	if err != nil {
		return err
	}
	out.Lines = f.Lines
	out.Header = f.Header
	out.Records = f.Records
	out.Trailer = f.Trailer
	return nil
}
