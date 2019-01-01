package gedcom5

func (r *IndividualRecord) Level() int {
	return r.lvl
}

func (r *IndividualRecord) SetLevel(lvl int) {
	r.lvl = lvl
}

func (r *PersonalName) Level() int {
	return r.lvl
}

func (r *PersonalName) SetLevel(lvl int) {
	r.lvl = lvl
}

func (r *Birth) Level() int {
	return r.lvl
}

func (r *Birth) SetLevel(lvl int) {
	r.lvl = lvl
}
