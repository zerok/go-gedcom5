//go:generate python3 generators/leveled.py --structs FamilyRecord --package gedcom5 --outfile record_fam_gen_leveled.go
//go:generate python3 generators/line_based.py --structs FamilyRecord --package gedcom5 --outfile record_fam_gen_linebased.go
package gedcom5

import "context"

type FamilyRecord struct {
	id    string
	lvl   int
	lines []Line

	Husband string `gedcom5:"HUSB"`
	Wife    string `gedcom5:"WIFE"`
}

func NewFamilyRecord() Record {
	return &FamilyRecord{
		lines: make([]Line, 0, 10),
	}
}

func (r *FamilyRecord) Decode(ctx context.Context) error {
	ld := NewLineDecoder(r, r.Level())
	return ld.Decode(ctx, r.Lines())
}
