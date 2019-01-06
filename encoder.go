package gedcom5

import (
	"context"
	"io"
	"reflect"

	"github.com/rs/zerolog"
)

type Encoder struct {
	err error
	out LineWriter
	ctx context.Context
}

func NewEncoder(out io.Writer) *Encoder {
	return &Encoder{
		out: newSimpleLineWriter(out),
		ctx: context.Background(),
	}
}

func (e *Encoder) Encode(f *File) error {
	e.writeLine(Line{Level: 0, Tag: "HEAD"})
	for _, l := range f.Header.Lines {
		e.writeLine(l)
	}
	for _, r := range f.Records {
		e.writeRecord(r)
	}
	e.writeLine(Line{Level: 0, Tag: "TRLR"})
	for _, l := range f.Trailer.Lines {
		e.writeLine(l)
	}
	return nil
}

func (e *Encoder) writeLine(line Line) error {
	if e.err != nil {
		return e.err
	}
	e.err = e.out.WriteLine(line)
	return e.err
}

func (e *Encoder) writeRecord(r Record) error {
	var tag string
	if e.err != nil {
		return e.err
	}
	switch r.(type) {
	case *IndividualRecord:
		tag = "INDI"
	}

	if tag != "" {
		l := Line{Tag: tag, Level: 0}
		if valuable, ok := r.(Valuable); ok {
			l.Value = valuable.Value()
		}
		e.writeLine(l)
		if e.err == nil {
			se := NewStructEncoder(r, 0)
			e.err = se.EncodeTo(e.ctx, e.out)
		}
	}
	return e.err
}

type StructEncoder struct {
	baseLevel int
	val       interface{}
}

func NewStructEncoder(val interface{}, lvl int) *StructEncoder {
	return &StructEncoder{
		baseLevel: lvl,
		val:       val,
	}
}

func (se *StructEncoder) processProperty(ctx context.Context, lw LineWriter, tag fieldTag, propValue interface{}) error {
	logger := zerolog.Ctx(ctx)
	val := ""
	typ := reflect.TypeOf(propValue)
	if s, ok := propValue.(string); ok {
		val = s
	} else if s, ok := propValue.(*string); ok {
		val = *s
	} else {
		if v, ok := propValue.(Valuable); ok {
			val = v.Value()
		}
	}
	isEmpty := val == ""
	blw := newBufferedLineWriter()
	if typ.Kind() == reflect.Struct || (typ.Kind() == reflect.Ptr && typ.Elem().Kind() == reflect.Struct) {
		subenc := NewStructEncoder(propValue, se.baseLevel+1)
		if err := subenc.EncodeTo(ctx, blw); err != nil {
			return err
		}
	}
	isEmpty = isEmpty && blw.IsEmpty()
	logger.Debug().Msgf("Tag: %s | Type: %s | Empty: %v\n", tag.Tag, typ.String(), isEmpty)
	if tag.OmitEmpty && isEmpty {
		return nil
	}
	if err := lw.WriteLine(Line{Tag: tag.Tag, Level: se.baseLevel + 1, Value: val}); err != nil {
		return err
	}
	return blw.WriteTo(lw)
}

func (se *StructEncoder) EncodeTo(ctx context.Context, lw LineWriter) error {
	typ := reflect.TypeOf(se.val)
	val := reflect.ValueOf(se.val)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return nil
	}
	for i := 0; i < typ.NumField(); i++ {
		field := typ.FieldByIndex([]int{i})
		tagValue := field.Tag.Get("gedcom5")
		if tagValue == "" {
			continue
		}
		tag := parseFieldTag(tagValue)
		if field.Type.Kind() == reflect.Slice {
			values := val.Elem().FieldByIndex([]int{i})
			for j := 0; j < values.Len(); j++ {
				value := values.Index(j)
				if err := se.processProperty(ctx, lw, tag, value.Addr().Interface()); err != nil {
					return err
				}
			}
		} else {
			propValue := val.Elem().FieldByIndex([]int{i}).Addr().Interface()
			if err := se.processProperty(ctx, lw, tag, propValue); err != nil {
				return err
			}
		}
	}
	return nil
}
