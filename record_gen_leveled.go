package gedcom5

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

func (r *IndividualRecord) Level() int {
	return r.lvl
}

func (r *IndividualRecord) SetLevel(lvl int) {
	r.lvl = lvl
}

func (r *FamilyRecord) Level() int {
	return r.lvl
}

func (r *FamilyRecord) SetLevel(lvl int) {
	r.lvl = lvl
}

func (r *MultimediaRecord) Level() int {
	return r.lvl
}

func (r *MultimediaRecord) SetLevel(lvl int) {
	r.lvl = lvl
}

func (r *RepositoryRecord) Level() int {
	return r.lvl
}

func (r *RepositoryRecord) SetLevel(lvl int) {
	r.lvl = lvl
}

func (r *NoteRecord) Level() int {
	return r.lvl
}

func (r *NoteRecord) SetLevel(lvl int) {
	r.lvl = lvl
}

func (r *SourceRecord) Level() int {
	return r.lvl
}

func (r *SourceRecord) SetLevel(lvl int) {
	r.lvl = lvl
}

func (r *SubmitterRecord) Level() int {
	return r.lvl
}

func (r *SubmitterRecord) SetLevel(lvl int) {
	r.lvl = lvl
}
