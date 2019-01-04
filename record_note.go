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
