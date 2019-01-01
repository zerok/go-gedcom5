package gedcom5

func (r *SourceRecord) Lines() []Line {
	return r.lines
}

func (r *SourceRecord) SetLines(lines []Line) {
	r.lines = lines
}

func (r *SourceRecord) AddLine(l Line) {
	if r.lines == nil {
		r.lines = make([]Line, 0, 10)
	}
	r.lines = append(r.lines, l)
}
