//go:generate python3 generators/leveled.py --structs RepositoryRecord --package gedcom5 --outfile record_repo_gen_leveled.go
//go:generate python3 generators/line_based.py --structs RepositoryRecord --package gedcom5 --outfile record_repo_gen_linebased.go
package gedcom5

import "context"

type RepositoryRecord struct {
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
