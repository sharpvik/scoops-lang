package primitives

import (
    "bufio"
    "fmt"
    "github.com/sharpvik/scoops/Package/Shared"
    "math"
)

type Float struct {
    Value float64
}


/*
func FloatToBytes(float float64) []byte {
    bits := math.Float64bits(float)
    bytes := make([]byte, 8)
    binary.LittleEndian.PutUint64(bytes, bits)
    return bytes
}
*/


func NewFloat(value float64) *Float {
    return &Float{value}
}


func (f *Float) Clone() shared.Object {
    return NewFloat(f.Value)
}


func (f *Float) Print(w *bufio.Writer) {
    w.WriteString( fmt.Sprintf("%v", f.Value) )
}


func (f *Float) Type() string {
    return "flt"
}


func AddFloat(a, b *Float) *Float {
    return NewFloat(a.Value + b.Value)
}


func SubFloat(a, b *Float) *Float {
    return NewFloat(a.Value - b.Value)
}


func MulFloat(a, b *Float) *Float {
    return NewFloat(a.Value * b.Value)
}


func DivFloat(a, b *Float) *Float {
    return NewFloat(a.Value / b.Value)
}


func PowFloat(a, b *Float) *Float {
    return NewFloat( math.Pow(a.Value, b.Value) )
}