package gedcom5

func (r *FamilyRecord) Level() int {
	return r.lvl
}

func (r *FamilyRecord) SetLevel(lvl int) {
	r.lvl = lvl
}
