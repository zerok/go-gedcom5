package gedcom5

func (r *IndividualRecord) ID() string {
	return r.id
}

func (r *IndividualRecord) SetID(id string) {
	r.id = id
}

func (r *FamilyRecord) ID() string {
	return r.id
}

func (r *FamilyRecord) SetID(id string) {
	r.id = id
}
