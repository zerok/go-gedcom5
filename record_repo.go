package gedcom5

type RepositoryRecord struct {
	lines []Line
}

func NewRepositoryRecord() Record {
	return &RepositoryRecord{
		lines: make([]Line, 0, 10),
	}
}

func (r *RepositoryRecord) AddLine(l Line) {
	r.lines = append(r.lines, l)
}

func (r *RepositoryRecord) Lines() []Line {
	return []Line{}
}
