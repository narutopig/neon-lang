package value

import (
	"bytes"
	"encoding/binary"
	"fmt"
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
	return Value{Type: Boolean, Data: boolToBytes(b)}
}

type ValueType byte

func floatToBytes(num float64) []byte {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, num)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	return buf.Bytes()
}

func boolToBytes(b bool) []byte {
	if b {
		return []byte{1}
	}

	return []byte{0}
}

const (
	String ValueType = iota
	Number
	Boolean
)
