package gedcom5

func (r *FamilyRecord) Lines() []Line {
	return r.lines
}

func (r *FamilyRecord) SetLines(lines []Line) {
	r.lines = lines
}

func (r *FamilyRecord) AddLine(l Line) {
	if r.lines == nil {
		r.lines = make([]Line, 0, 10)
	}
	r.lines = append(r.lines, l)
}
