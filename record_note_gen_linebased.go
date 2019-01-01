package gedcom5

func (r *NoteRecord) Lines() []Line {
	return r.lines
}

func (r *NoteRecord) SetLines(lines []Line) {
	r.lines = lines
}

func (r *NoteRecord) AddLine(l Line) {
	if r.lines == nil {
		r.lines = make([]Line, 0, 10)
	}
	r.lines = append(r.lines, l)
}
