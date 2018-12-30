package gedcom5

type SubmitterRecord struct {
	lines []Line
}

func NewSubmitterRecord() Record {
	return &SubmitterRecord{
		lines: make([]Line, 0, 10),
	}
}

func (r *SubmitterRecord) AddLine(l Line) {
	r.lines = append(r.lines, l)
}

func (r *SubmitterRecord) Lines() []Line {
	return []Line{}
}
