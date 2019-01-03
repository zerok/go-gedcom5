//go:generate python3 generators/leveled.py --structs Event --package gedcom5 --outfile event_gen_leveled.go
//go:generate python3 generators/line_based.py --structs Event --package gedcom5 --outfile event_gen_linebased.go
package gedcom5

import "context"

type Event struct {
	lvl   int
	lines []Line

	Address Address  `gedcom5:"ADDR"`
	Phones  []string `gedcom5:"PHONE"`
}

func (e *Event) String() string {
	return e.Address.Value()
}

func (e *Event) Decode(ctx context.Context) error {
	ld := NewLineDecoder(e, e.lvl)
	return ld.Decode(ctx, e.lines)
}

type Address struct {
	value string
}

func (a *Address) Value() string {
	return a.value
}
func (a *Address) SetValue(val string) {
	a.value = val
}