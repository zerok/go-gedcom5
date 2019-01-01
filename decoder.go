package gedcom5

import (
	"context"
	"io"
)

type Decoder struct {
	p   *FileParser
	ctx context.Context
}

func (d *Decoder) WithContext(ctx context.Context) *Decoder {
	return &Decoder{
		p:   d.p,
		ctx: ctx,
	}
}

func NewDecoder(in io.Reader) *Decoder {
	s := NewScanner(in)
	return &Decoder{
		p:   NewFileParser(s),
		ctx: context.Background(),
	}
}

func (ld *Decoder) Decode(out *File) error {
	f, err := ld.p.ParseFile(ld.ctx)
	if err != nil {
		return err
	}
	out.Lines = f.Lines
	out.Header = f.Header
	out.Records = f.Records
	out.Trailer = f.Trailer
	return nil
}
