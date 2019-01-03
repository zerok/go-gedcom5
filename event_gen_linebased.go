package gedcom5

func (r *Event) Lines() []Line {
	return r.lines
}

func (r *Event) SetLines(lines []Line) {
	r.lines = lines
}

func (r *Event) AddLine(l Line) {
	if r.lines == nil {
		r.lines = make([]Line, 0, 10)
	}
	r.lines = append(r.lines, l)
}
