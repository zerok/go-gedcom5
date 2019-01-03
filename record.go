//go:generate python3 generators/identifyable.py --outfile record_identifyable.go --package gedcom5 --struct IndividualRecord FamilyRecord
package gedcom5

import "context"

type Lined interface {
	AddLine(Line)
	Lines() []Line
	SetLines([]Line)
}

type Identifyable interface {
	ID() string
	SetID(string)
}

type Decodable interface {
	Decode(context.Context) error
}

type Record interface {
	Lined
	Decodable
}

type Leveled interface {
	SetLevel(int)
	Level() int
}

type Valuable interface {
	SetValue(string)
	Value() string
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

func (r *UnknownRecord) Decode(ctx context.Context) error {
	return nil
}

func (r *UnknownRecord) AddLine(l Line) {
}

func (r *UnknownRecord) Lines() []Line {
	return []Line{}
}

func (r *UnknownRecord) SetLines(l []Line) {
}
