package gedcom5

type MultimediaRecord struct {
	lines []Line
}

func NewMultimediaRecord() Record {
	return &MultimediaRecord{
		lines: make([]Line, 0, 10),
	}
}

func (r *MultimediaRecord) AddLine(l Line) {
	r.lines = append(r.lines, l)
}

func (r *MultimediaRecord) Lines() []Line {
	return []Line{}
}
