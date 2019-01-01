package gedcom5

func (r *IndividualRecord) Lines() []Line {
	return r.lines
}

func (r *IndividualRecord) SetLines(lines []Line) {
	r.lines = lines
}

func (r *IndividualRecord) AddLine(l Line) {
	if r.lines == nil {
		r.lines = make([]Line, 0, 10)
	}
	r.lines = append(r.lines, l)
}

func (r *Birth) Lines() []Line {
	return r.lines
}

func (r *Birth) SetLines(lines []Line) {
	r.lines = lines
}

func (r *Birth) AddLine(l Line) {
	if r.lines == nil {
		r.lines = make([]Line, 0, 10)
	}
	r.lines = append(r.lines, l)
}

func (r *PersonalName) Lines() []Line {
	return r.lines
}

func (r *PersonalName) SetLines(lines []Line) {
	r.lines = lines
}

func (r *PersonalName) AddLine(l Line) {
	if r.lines == nil {
		r.lines = make([]Line, 0, 10)
	}
	r.lines = append(r.lines, l)
}
