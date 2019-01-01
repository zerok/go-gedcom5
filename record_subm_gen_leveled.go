package gedcom5

func (r *SubmitterRecord) Level() int {
	return r.lvl
}

func (r *SubmitterRecord) SetLevel(lvl int) {
	r.lvl = lvl
}
