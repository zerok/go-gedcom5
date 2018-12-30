package gedcom5

type IndividualRecord struct {
	lines []Line
}

func NewIndividualRecord() Record {
	return &IndividualRecord{
		lines: make([]Line, 0, 10),
	}
}

func (r *IndividualRecord) AddLine(l Line) {
	r.lines = append(r.lines, l)
}

func (r *IndividualRecord) Lines() []Line {
	return []Line{}
}
