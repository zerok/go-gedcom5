package gedcom5

type FamilyRecord struct {
	lines []Line
}

func NewFamilyRecord() Record {
	return &FamilyRecord{
		lines: make([]Line, 0, 10),
	}
}

func (r *FamilyRecord) AddLine(l Line) {
	r.lines = append(r.lines, l)
}

func (r *FamilyRecord) Lines() []Line {
	return []Line{}
}
