// Code generated by github.com/actgardner/gogen-avro/v8. DO NOT EDIT.
/*
 * SOURCE:
 *     evolution.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v8/compiler"
	"github.com/actgardner/gogen-avro/v8/vm"
	"github.com/actgardner/gogen-avro/v8/vm/types"
)

type UnionStringLongIntFloatDoubleNullTypeEnum int

const (
	UnionStringLongIntFloatDoubleNullTypeEnumString UnionStringLongIntFloatDoubleNullTypeEnum = 0

	UnionStringLongIntFloatDoubleNullTypeEnumLong UnionStringLongIntFloatDoubleNullTypeEnum = 1

	UnionStringLongIntFloatDoubleNullTypeEnumInt UnionStringLongIntFloatDoubleNullTypeEnum = 2

	UnionStringLongIntFloatDoubleNullTypeEnumFloat UnionStringLongIntFloatDoubleNullTypeEnum = 3

	UnionStringLongIntFloatDoubleNullTypeEnumDouble UnionStringLongIntFloatDoubleNullTypeEnum = 4
)

type UnionStringLongIntFloatDoubleNull struct {
	String    string
	Long      int64
	Int       int32
	Float     float32
	Double    float64
	Null      *types.NullVal
	UnionType UnionStringLongIntFloatDoubleNullTypeEnum
}

func writeUnionStringLongIntFloatDoubleNull(r *UnionStringLongIntFloatDoubleNull, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(5, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionStringLongIntFloatDoubleNullTypeEnumString:
		return vm.WriteString(r.String, w)
	case UnionStringLongIntFloatDoubleNullTypeEnumLong:
		return vm.WriteLong(r.Long, w)
	case UnionStringLongIntFloatDoubleNullTypeEnumInt:
		return vm.WriteInt(r.Int, w)
	case UnionStringLongIntFloatDoubleNullTypeEnumFloat:
		return vm.WriteFloat(r.Float, w)
	case UnionStringLongIntFloatDoubleNullTypeEnumDouble:
		return vm.WriteDouble(r.Double, w)
	}
	return fmt.Errorf("invalid value for *UnionStringLongIntFloatDoubleNull")
}

func NewUnionStringLongIntFloatDoubleNull() *UnionStringLongIntFloatDoubleNull {
	return &UnionStringLongIntFloatDoubleNull{}
}

func (r *UnionStringLongIntFloatDoubleNull) Serialize(w io.Writer) error {
	return writeUnionStringLongIntFloatDoubleNull(r, w)
}

func DeserializeUnionStringLongIntFloatDoubleNull(r io.Reader) (*UnionStringLongIntFloatDoubleNull, error) {
	t := NewUnionStringLongIntFloatDoubleNull()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func (r *UnionStringLongIntFloatDoubleNull) Schema() string {
	return "[\"string\",\"long\",\"int\",\"float\",\"double\",\"null\"]"
}

func (_ *UnionStringLongIntFloatDoubleNull) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionStringLongIntFloatDoubleNull) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionStringLongIntFloatDoubleNull) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionStringLongIntFloatDoubleNull) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionStringLongIntFloatDoubleNull) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionStringLongIntFloatDoubleNull) SetString(v string)  { panic("Unsupported operation") }
func (r *UnionStringLongIntFloatDoubleNull) SetLong(v int64) {
	r.UnionType = (UnionStringLongIntFloatDoubleNullTypeEnum)(v)
}
func (r *UnionStringLongIntFloatDoubleNull) Get(i int) types.Field {
	switch i {
	case 0:
		return &types.String{Target: (&r.String)}
	case 1:
		return &types.Long{Target: (&r.Long)}
	case 2:
		return &types.Int{Target: (&r.Int)}
	case 3:
		return &types.Float{Target: (&r.Float)}
	case 4:
		return &types.Double{Target: (&r.Double)}
	case 5:
		return r.Null
	}
	panic("Unknown field index")
}
func (_ *UnionStringLongIntFloatDoubleNull) NullField(i int)  { panic("Unsupported operation") }
func (_ *UnionStringLongIntFloatDoubleNull) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionStringLongIntFloatDoubleNull) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ *UnionStringLongIntFloatDoubleNull) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionStringLongIntFloatDoubleNull) Finalize()                {}

func (r *UnionStringLongIntFloatDoubleNull) MarshalJSON() ([]byte, error) {
	if r == nil {
		return []byte("null"), nil
	}
	switch r.UnionType {
	case UnionStringLongIntFloatDoubleNullTypeEnumString:
		return json.Marshal(map[string]interface{}{"string": r.String})
	case UnionStringLongIntFloatDoubleNullTypeEnumLong:
		return json.Marshal(map[string]interface{}{"long": r.Long})
	case UnionStringLongIntFloatDoubleNullTypeEnumInt:
		return json.Marshal(map[string]interface{}{"int": r.Int})
	case UnionStringLongIntFloatDoubleNullTypeEnumFloat:
		return json.Marshal(map[string]interface{}{"float": r.Float})
	case UnionStringLongIntFloatDoubleNullTypeEnumDouble:
		return json.Marshal(map[string]interface{}{"double": r.Double})
	}
	return nil, fmt.Errorf("invalid value for *UnionStringLongIntFloatDoubleNull")
}

func (r *UnionStringLongIntFloatDoubleNull) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if value, ok := fields["string"]; ok {
		r.UnionType = 0
		return json.Unmarshal([]byte(value), &r.String)
	}
	if value, ok := fields["long"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Long)
	}
	if value, ok := fields["int"]; ok {
		r.UnionType = 2
		return json.Unmarshal([]byte(value), &r.Int)
	}
	if value, ok := fields["float"]; ok {
		r.UnionType = 3
		return json.Unmarshal([]byte(value), &r.Float)
	}
	if value, ok := fields["double"]; ok {
		r.UnionType = 4
		return json.Unmarshal([]byte(value), &r.Double)
	}
	return fmt.Errorf("invalid value for *UnionStringLongIntFloatDoubleNull")
}
