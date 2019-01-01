//go:generate python3 generators/leveled.py --structs SubmitterRecord --package gedcom5 --outfile record_subm_gen_leveled.go
//go:generate python3 generators/line_based.py --structs SubmitterRecord --package gedcom5 --outfile record_subm_gen_linebased.go
package gedcom5

import "context"

type SubmitterRecord struct {
	lvl   int
	lines []Line
}

func NewSubmitterRecord() Record {
	return &SubmitterRecord{
		lines: make([]Line, 0, 10),
	}
}

func (r *SubmitterRecord) Decode(ctx context.Context) error {
	return nil
}
