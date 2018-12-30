package gedcom5

type Record interface {
	AddLine(Line)
	Lines() []Line
}

func NewRecord(tag string) Record {
	switch tag {
	case "INDI":
		return NewIndividualRecord()
	case "FAM":
		return NewFamilyRecord()
	case "OBJE":
		return NewMultimediaRecord()
	case "NOTE":
		return NewNoteRecord()
	case "REPO":
		return NewRepositoryRecord()
	case "SOUR":
		return NewSourceRecord()
	case "SUBM":
		return NewSubmitterRecord()
	default:
		return &UnknownRecord{}
	}
}

type UnknownRecord struct {
}

func (r *UnknownRecord) AddLine(l Line) {
}

func (r *UnknownRecord) Lines() []Line {
	return []Line{}
}
