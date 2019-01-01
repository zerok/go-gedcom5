package gedcom5

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
