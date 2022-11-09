package value

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

type Value struct {
	Type ValueType
	Data []byte
}

func NewString(str string) Value {
	return Value{Type: String, Data: []byte(str)}
}

func NewNumber(num float64) Value {
	return Value{Type: Number, Data: floatToBytes(num)}
}

func NewBoolean(b bool) Value {
	return Value{Type: Boolean, Data: BoolToBytes(b)}
}

func NewIdentifier(id string) Value {
	return Value{Type: Identifier, Data: []byte(id)}
}

func (v Value) String() string {
	var data any
	switch v.Type {
	case String:
		data = string(v.Data)
	case Number:
		data = FloatFromBytes(v.Data)
	case Boolean:
		if v.Data[0] == 0 {
			data = False
		} else {
			data = True
		}
	}

	return fmt.Sprintf("Value{Type: %s, Value: %s}", v.Type, data)
}

// Value enum
const (
	String ValueType = iota
	Number
	Boolean
	Identifier
)

// Boolean constants
var (
	True  = NewBoolean(true)
	False = NewBoolean(false)
)

// ValueType describes what type a value is
//
//go:generate stringer -type=ValueType
type ValueType byte

func floatToBytes(num float64) []byte {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, num)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	return buf.Bytes()
}

func BoolToBytes(b bool) []byte {
	if b {
		return []byte{1}
	}

	return []byte{0}
}

func FloatFromBytes(bytes []byte) float64 {
	return math.Float64frombits(binary.BigEndian.Uint64(bytes))
}
