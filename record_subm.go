package gedcom5

import "context"

type SubmitterRecord struct {
	id    string
	lvl   int
	lines []Line
}

func NewSubmitterRecord() Record {
	return &SubmitterRecord{
		lines: make([]Line, 0, 10),
	}
}

func (r *SubmitterRecord) Decode(ctx context.Context) error {
	return nil
}
