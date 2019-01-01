package gedcom5

func (r *RepositoryRecord) Level() int {
	return r.lvl
}

func (r *RepositoryRecord) SetLevel(lvl int) {
	r.lvl = lvl
}
