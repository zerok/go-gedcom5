//go:generate python3 generators/leveled.py --structs NoteRecord --package gedcom5 --outfile record_note_gen_leveled.go
//go:generate python3 generators/line_based.py --structs NoteRecord --package gedcom5 --outfile record_note_gen_linebased.go
package gedcom5

import "context"

type NoteRecord struct {
	id    string
	lvl   int
	lines []Line
}

func NewNoteRecord() Record {
	return &NoteRecord{
		lines: make([]Line, 0, 10),
	}
}

func (r *NoteRecord) Decode(ctx context.Context) error {
	return nil
}
