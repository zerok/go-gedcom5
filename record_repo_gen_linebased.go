package gedcom5

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
