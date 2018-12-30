package gedcom5

type SourceRecord struct {
	lines []Line
}

func NewSourceRecord() Record {
	return &SourceRecord{
		lines: make([]Line, 0, 10),
	}
}

func (r *SourceRecord) AddLine(l Line) {
	r.lines = append(r.lines, l)
}

func (r *SourceRecord) Lines() []Line {
	return []Line{}
}
