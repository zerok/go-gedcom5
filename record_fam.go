package gedcom5

import "context"

type FamilyRecord struct {
	id    string
	lvl   int
	lines []Line

	Husband  string   `gedcom5:"HUSB"`
	Wife     string   `gedcom5:"WIFE"`
	Children []string `gedcom5:"CHIL"`
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
