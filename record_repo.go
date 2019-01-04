package gedcom5

import "context"

type RepositoryRecord struct {
	id    string
	lvl   int
	lines []Line
}

func NewRepositoryRecord() Record {
	return &RepositoryRecord{
		lines: make([]Line, 0, 10),
	}
}

func (r *RepositoryRecord) Decode(ctx context.Context) error {
	return nil
}
