package gedcom5

type NoteRecord struct {
	lines []Line
}

func NewNoteRecord() Record {
	return &NoteRecord{
		lines: make([]Line, 0, 10),
	}
}

func (r *NoteRecord) AddLine(l Line) {
	r.lines = append(r.lines, l)
}

func (r *NoteRecord) Lines() []Line {
	return []Line{}
}
