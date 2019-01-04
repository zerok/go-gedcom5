package gedcom5

import "context"

type MultimediaRecord struct {
	id    string
	lvl   int
	lines []Line
}

func NewMultimediaRecord() Record {
	return &MultimediaRecord{
		lines: make([]Line, 0, 10),
	}
}

func (r *MultimediaRecord) Decode(ctx context.Context) error {
	return nil
}
