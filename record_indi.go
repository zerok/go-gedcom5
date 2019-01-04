package gedcom5

import (
	"bytes"
	"context"
)

const structTagName = "gedcom5"

type IndividualRecord struct {
	id    string
	lvl   int
	lines []Line

	PersonalNames []PersonalName `gedcom5:"NAME,omitempty"`
	Sex           string         `gedcom5:"SEX,omitempty"`
	Religion      string         `gedcom5:"RELI,omitempty"`
	Note          string         `gedcom5:"NOTE,omitempty"`
	Occupation    string         `gedcom5:"OCCU,omitempty"`
	Birth         Birth          `gedcom5:"BIRT,omitempty"`
	Residence     Event          `gedcom5:"RESI,omitempty"`
}

func (r *IndividualRecord) String() string {
	if r.PersonalNames == nil || len(r.PersonalNames) == 0 {
		return "<no name>"
	}
	return r.PersonalNames[0].String()
}

func NewIndividualRecord() Record {
	return &IndividualRecord{
		lines: make([]Line, 0, 10),
	}
}

func (r *IndividualRecord) Decode(ctx context.Context) error {
	ld := NewLineDecoder(r, r.Level())
	return ld.Decode(ctx, r.Lines())
}

type PersonalName struct {
	lvl           int
	lines         []Line
	Name          string
	Prefix        string `gedcom5:"NPFX,omitempty"`
	Suffix        string `gedcom5:"NSFX,omitempty"`
	Given         string `gedcom5:"GIVN,omitempty"`
	SurnamePrefix string `gedcom5:"SPFX,omitempty"`
	Surname       string `gedcom5:"SURN,omitempty"`
}

func (pn *PersonalName) SetValue(v string) {
	pn.Name = v
}

func (pn *PersonalName) Value() string {
	return pn.Name
}

func (pn *PersonalName) String() string {
	if pn.Name != "" {
		return pn.Name
	}
	empty := true
	var out bytes.Buffer
	elems := []string{pn.Prefix, pn.Given, pn.SurnamePrefix, pn.Surname, pn.Suffix}
	for _, elem := range elems {
		if elem != "" {
			if !empty {
				out.WriteByte(0x20)
			}
			out.WriteString(elem)
			empty = false
		}
	}
	return out.String()
}

func (pn *PersonalName) Decode(ctx context.Context) error {
	ld := NewLineDecoder(pn, pn.Level())
	return ld.Decode(ctx, pn.Lines())
}

type Birth struct {
	lvl   int
	lines []Line

	Date string `gedcom5:"DATE"`
}

func (b *Birth) Decode(ctx context.Context) error {
	ld := NewLineDecoder(b, b.lvl)
	return ld.Decode(ctx, b.Lines())
}
