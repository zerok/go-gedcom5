package gedcom5

import "context"

type NoteRecord struct {
	id    string
	lvl   int
	lines []Line
	value string
}

func NewNoteRecord() Record {
	return &NoteRecord{
		lines: make([]Line, 0, 10),
	}
}

func (r *NoteRecord) Decode(ctx context.Context) error {
	ld := NewLineDecoder(r, r.Level())
	return ld.Decode(ctx, r.Lines())
}

func (r *NoteRecord) Value() string {
	return r.value
}

func (r *NoteRecord) SetValue(v string) {
	r.value = v
}
