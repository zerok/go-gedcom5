package gedcom5

func (r *SourceRecord) Level() int {
	return r.lvl
}

func (r *SourceRecord) SetLevel(lvl int) {
	r.lvl = lvl
}
