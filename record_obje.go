//go:generate python3 generators/leveled.py --structs MultimediaRecord --package gedcom5 --outfile record_obje_gen_leveled.go
//go:generate python3 generators/line_based.py --structs MultimediaRecord --package gedcom5 --outfile record_obje_gen_linebased.go
package gedcom5

import "context"

type MultimediaRecord struct {
	id    string
	lvl   int
	lines []Line
}

func NewMultimediaRecord() Record {
	return &MultimediaRecord{
		lines: make([]Line, 0, 10),
	}
}

func (r *MultimediaRecord) Decode(ctx context.Context) error {
	return nil
}
