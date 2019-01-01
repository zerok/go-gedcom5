package gedcom5

func (r *MultimediaRecord) Level() int {
	return r.lvl
}

func (r *MultimediaRecord) SetLevel(lvl int) {
	r.lvl = lvl
}
