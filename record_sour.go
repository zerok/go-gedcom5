package gedcom5

import "context"

type SourceRecord struct {
	lvl   int
	lines []Line
}

func NewSourceRecord() Record {
	return &SourceRecord{
		lines: make([]Line, 0, 10),
	}
}

func (r *SourceRecord) Decode(ctx context.Context) error {
	return nil
}
