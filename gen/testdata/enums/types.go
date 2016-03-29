// Code generated by thriftrw

package enums

import (
	"fmt"
	"github.com/thriftrw/thriftrw-go/wire"
	"strings"
)

type EmptyEnum int32

const ()

func (v EmptyEnum) ToWire() wire.Value {
	return wire.NewValueI32(int32(v))
}
func (v *EmptyEnum) FromWire(w wire.Value) error {
	*v = (EmptyEnum)(w.GetI32())
	return nil
}

type EnumDefault int32

const (
	EnumDefaultFoo EnumDefault = 0
	EnumDefaultBar EnumDefault = 1
	EnumDefaultBaz EnumDefault = 2
)

func (v EnumDefault) ToWire() wire.Value {
	return wire.NewValueI32(int32(v))
}
func (v *EnumDefault) FromWire(w wire.Value) error {
	*v = (EnumDefault)(w.GetI32())
	return nil
}

type EnumWithDuplicateName int32

const (
	EnumWithDuplicateNameA EnumWithDuplicateName = 0
	EnumWithDuplicateNameB EnumWithDuplicateName = 1
	EnumWithDuplicateNameC EnumWithDuplicateName = 2
	EnumWithDuplicateNameP EnumWithDuplicateName = 3
	EnumWithDuplicateNameQ EnumWithDuplicateName = 4
	EnumWithDuplicateNameR EnumWithDuplicateName = 5
	EnumWithDuplicateNameX EnumWithDuplicateName = 6
	EnumWithDuplicateNameY EnumWithDuplicateName = 7
	EnumWithDuplicateNameZ EnumWithDuplicateName = 8
)

func (v EnumWithDuplicateName) ToWire() wire.Value {
	return wire.NewValueI32(int32(v))
}
func (v *EnumWithDuplicateName) FromWire(w wire.Value) error {
	*v = (EnumWithDuplicateName)(w.GetI32())
	return nil
}

type EnumWithDuplicateValues int32

const (
	EnumWithDuplicateValuesP EnumWithDuplicateValues = 0
	EnumWithDuplicateValuesQ EnumWithDuplicateValues = -1
	EnumWithDuplicateValuesR EnumWithDuplicateValues = 0
)

func (v EnumWithDuplicateValues) ToWire() wire.Value {
	return wire.NewValueI32(int32(v))
}
func (v *EnumWithDuplicateValues) FromWire(w wire.Value) error {
	*v = (EnumWithDuplicateValues)(w.GetI32())
	return nil
}

type EnumWithValues int32

const (
	EnumWithValuesX EnumWithValues = 123
	EnumWithValuesY EnumWithValues = 456
	EnumWithValuesZ EnumWithValues = 789
)

func (v EnumWithValues) ToWire() wire.Value {
	return wire.NewValueI32(int32(v))
}
func (v *EnumWithValues) FromWire(w wire.Value) error {
	*v = (EnumWithValues)(w.GetI32())
	return nil
}

type StructWithOptionalEnum struct{ E *EnumDefault }

func (v *StructWithOptionalEnum) ToWire() wire.Value {
	var fields [1]wire.Field
	i := 0
	if v.E != nil {
		fields[i] = wire.Field{ID: 1, Value: v.E.ToWire()}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]})
}
func _EnumDefault_Read(w wire.Value) (EnumDefault, error) {
	var v EnumDefault
	err := v.FromWire(w)
	return v, err
}
func (v *StructWithOptionalEnum) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TI32 {
				var x EnumDefault
				x, err = _EnumDefault_Read(field.Value)
				v.E = &x
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
func (v *StructWithOptionalEnum) String() string {
	var fields [1]string
	i := 0
	if v.E != nil {
		fields[i] = fmt.Sprintf("E: %v", *(v.E))
		i++
	}
	return fmt.Sprintf("StructWithOptionalEnum{%v}", strings.Join(fields[:i], ", "))
}
