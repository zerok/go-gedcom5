package gedcom5

import (
	"bytes"
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type Line struct {
	Level  int
	XRefID string
	Tag    string
	Value  string
}

func (l *Line) String() string {
	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("%d", l.Level))
	if l.XRefID != "" {
		out.WriteByte(0x20)
		out.WriteString(l.XRefID)
	}
	if l.Tag != "" {
		out.WriteByte(0x20)
		out.WriteString(l.Tag)
	}
	if l.Value != "" {
		out.WriteByte(0x20)
		out.WriteString(l.Value)
	}
	return out.String()
}

// LineDecoder handles reflection-based decoding of multiple gedcom5 Lines into
// a single struct.
type LineDecoder struct {
	baseLevel              int
	val                    interface{}
	tagToField             map[string]string
	tagToType              map[string]reflect.Type
	tagValue               string
	previousStructField    string
	previousPrimitiveField string
	pendingLines           []Line
}

func NewLineDecoder(val interface{}, baseLevel int) *LineDecoder {
	ld := &LineDecoder{
		val:        val,
		baseLevel:  baseLevel,
		tagToField: make(map[string]string),
		tagToType:  make(map[string]reflect.Type),
	}
	return ld
}

func (ld *LineDecoder) buildMappings() error {
	typ := reflect.ValueOf(ld.val).Elem().Type()
	fieldCount := typ.NumField()
	for i := 0; i < fieldCount; i++ {
		field := typ.Field(i)
		tagValue := field.Tag.Get(structTagName)
		if tagValue != "" {
			tag := strings.Split(tagValue, ",")[0]
			ld.tagToField[tag] = field.Name
			ld.tagToType[tag] = field.Type
		}
	}
	return nil
}

func (ld *LineDecoder) decodePreviousField(ctx context.Context) error {
	field := ld.tagToField[ld.previousStructField]
	container := reflect.ValueOf(ld.val).Elem()
	prop := container.FieldByName(field)
	var i interface{}
	if prop.Kind() == reflect.Slice {
		i = reflect.New(prop.Type().Elem()).Interface()
	} else {
		i = prop.Addr().Interface()
	}
	if lvld, ok := i.(Leveled); ok {
		lvld.SetLevel(ld.baseLevel + 1)
	}
	if lined, ok := i.(Lined); ok {
		lined.SetLines(ld.pendingLines)
	}
	if ld.tagValue != "" {
		if v, ok := i.(Valuable); ok {
			v.SetValue(ld.tagValue)
		}
	}
	if decodable, ok := i.(Decodable); ok {
		if err := decodable.Decode(ctx); err != nil {
			return err
		}
	}
	if prop.Kind() == reflect.Slice {
		rslice := reflect.Append(prop, reflect.Indirect(reflect.ValueOf(i)))
		prop.Set(rslice)
	}
	return nil
}

func (ld *LineDecoder) resetFieldData(tag string) {
	ld.previousStructField = tag
	ld.pendingLines = make([]Line, 0, 10)
	ld.tagValue = ""
}

func (ld *LineDecoder) Decode(ctx context.Context, lines []Line) error {
	logger := zerolog.Ctx(ctx)
	if err := ld.buildMappings(); err != nil {
		return err
	}
	for _, line := range lines {
		if line.Level == ld.baseLevel+1 {
			field, ok := ld.tagToField[line.Tag]
			if !ok {
				logger.Debug().Msgf("No field found for tag %s", line.Tag)
				continue
			}
			typ, ok := ld.tagToType[line.Tag]
			if !ok {
				logger.Debug().Msgf("No type found for tag %s", line.Tag)
				continue
			}
			switch typ.Kind() {
			case reflect.String:
				if ld.previousStructField != "" {
					if err := ld.decodePreviousField(ctx); err != nil {
						return errors.Wrapf(err, "failed to decode tag %s", ld.previousStructField)
					}
				}
				ld.resetFieldData("")
				reflect.ValueOf(ld.val).Elem().FieldByName(field).SetString(line.Value)
				ld.previousPrimitiveField = field
			case reflect.Struct:
				if ld.previousStructField != "" {
					if err := ld.decodePreviousField(ctx); err != nil {
						return errors.Wrapf(err, "failed to decode tag %s", ld.previousStructField)
					}
				}
				ld.resetFieldData(line.Tag)
				if line.Value != "" {
					ld.tagValue = line.Value
				}
			case reflect.Slice:
				if ld.previousStructField != "" {
					if err := ld.decodePreviousField(ctx); err != nil {
						return errors.Wrapf(err, "failed to decode tag %s", ld.previousStructField)
					}
				}
				ld.resetFieldData(line.Tag)
				if line.Value != "" {
					ld.tagValue = line.Value
				}
			default:
				logger.Warn().Msgf("Unsupported kind of tag %s: %s", line.Tag, typ.Kind())
				ld.resetFieldData(line.Tag)
			}
		} else if line.Level > ld.baseLevel+1 {
			if line.Level == ld.baseLevel+2 && ld.previousPrimitiveField != "" {
				prevValue := reflect.ValueOf(ld.val).Elem().FieldByName(ld.previousPrimitiveField).Interface().(string)
				newValue := ""
				switch line.Tag {
				case "CONT":
					newValue = prevValue + "\n" + line.Value
				case "CONC":
					newValue = prevValue + line.Value
				}
				if newValue != "" {
					reflect.ValueOf(ld.val).Elem().FieldByName(ld.previousPrimitiveField).SetString(newValue)
				}

			}
			ld.pendingLines = append(ld.pendingLines, line)
		}
	}
	if ld.previousStructField != "" {
		if err := ld.decodePreviousField(ctx); err != nil {
			return errors.Wrapf(err, "failed to decode tag %s", ld.previousStructField)
		}
	}
	return nil
}
