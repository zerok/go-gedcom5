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

func (r *MultimediaRecord) Lines() []Line {
	return r.lines
}

func (r *MultimediaRecord) SetLines(lines []Line) {
	r.lines = lines
}

func (r *MultimediaRecord) AddLine(l Line) {
	if r.lines == nil {
		r.lines = make([]Line, 0, 10)
	}
	r.lines = append(r.lines, l)
}

func (r *RepositoryRecord) Lines() []Line {
	return r.lines
}

func (r *RepositoryRecord) SetLines(lines []Line) {
	r.lines = lines
}

func (r *RepositoryRecord) AddLine(l Line) {
	if r.lines == nil {
		r.lines = make([]Line, 0, 10)
	}
	r.lines = append(r.lines, l)
}

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
