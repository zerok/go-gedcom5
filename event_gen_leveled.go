package gedcom5

func (r *Event) Level() int {
	return r.lvl
}

func (r *Event) SetLevel(lvl int) {
	r.lvl = lvl
}
