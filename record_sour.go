package gedcom5

import "context"

type SourceRecord struct {
	lvl   int
	lines []Line

	Title string `gedcom5:"TITL"`
}

func NewSourceRecord() Record {
	return &SourceRecord{
		lines: make([]Line, 0, 10),
	}
}

func (r *SourceRecord) Decode(ctx context.Context) error {
	ld := NewLineDecoder(r, r.Level())
	return ld.Decode(ctx, r.Lines())
}
