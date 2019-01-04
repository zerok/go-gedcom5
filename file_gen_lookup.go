package gedcom5

func (f *File) LookupIndividualByID(id string) (*IndividualRecord, bool) {
	for _, r := range f.Records {
		t, tok := r.(*IndividualRecord)
		i, iok := r.(Identifyable)
		if tok && iok && i.ID() == id {
			return t, true
		}
	}
	return nil, false
}

func (f *File) LookupFamilyByID(id string) (*FamilyRecord, bool) {
	for _, r := range f.Records {
		t, tok := r.(*FamilyRecord)
		i, iok := r.(Identifyable)
		if tok && iok && i.ID() == id {
			return t, true
		}
	}
	return nil, false
}
