package gedcom5

func (r *SubmitterRecord) Lines() []Line {
	return r.lines
}

func (r *SubmitterRecord) SetLines(lines []Line) {
	r.lines = lines
}

func (r *SubmitterRecord) AddLine(l Line) {
	if r.lines == nil {
		r.lines = make([]Line, 0, 10)
	}
	r.lines = append(r.lines, l)
}
