//go:generate python3 generators/leveled.py --structs FamilyRecord --package gedcom5 --outfile record_fam_gen_leveled.go
//go:generate python3 generators/line_based.py --structs FamilyRecord --package gedcom5 --outfile record_fam_gen_linebased.go
package gedcom5

import "context"

type FamilyRecord struct {
	lvl   int
	lines []Line
}

func NewFamilyRecord() Record {
	return &FamilyRecord{
		lines: make([]Line, 0, 10),
	}
}

func (r *FamilyRecord) Decode(ctx context.Context) error {
	return nil
}
