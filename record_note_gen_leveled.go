package gedcom5

func (r *NoteRecord) Level() int {
	return r.lvl
}

func (r *NoteRecord) SetLevel(lvl int) {
	r.lvl = lvl
}
