//go:generate python3 generators/leveled.py --structs SourceRecord --package gedcom5 --outfile record_sour_gen_leveled.go
//go:generate python3 generators/line_based.py --structs SourceRecord --package gedcom5 --outfile record_sour_gen_linebased.go
package gedcom5

import "context"

type SourceRecord struct {
	lvl   int
	lines []Line
}

func NewSourceRecord() Record {
	return &SourceRecord{
		lines: make([]Line, 0, 10),
	}
}

func (r *SourceRecord) Decode(ctx context.Context) error {
	return nil
}
