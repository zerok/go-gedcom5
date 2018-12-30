package gedcom5

import "io"

type Decoder struct {
	p *LineParser
}

func NewDecoder(in io.Reader) *Decoder {
	return &Decoder{p: NewLineParser(in)}
}

func (ld *Decoder) Decode(out *File) error {
	out.Lines = make([]Line, 0, 10)
	for {
		l, err := ld.p.ParseLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		out.Lines = append(out.Lines, *l)
	}
	return nil
}
