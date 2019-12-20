package primitives

import (
    "bufio"
    "fmt"
    "github.com/sharpvik/scoops/Package/Shared"
)



type Byte struct {
    Value byte
}


func NewByte(value byte) *Byte {
    return &Byte{value}
}


func (b *Byte) Clone() shared.Object {
    return NewByte(b.Value)
}


func (b *Byte) Print(w *bufio.Writer) {
    w.WriteString( fmt.Sprintf("%d", b.Value) )
    w.Flush()
}


func (b *Byte) Type() string {
    return "byte"
}


func AddByte(a, b *Byte) *Byte {
    return NewByte(a.Value + b.Value)
}


func SubByte(a, b *Byte) *Byte {
    return NewByte(a.Value - b.Value)
}


func MulByte(a, b *Byte) *Byte {
    return NewByte(a.Value * b.Value)
}


func DivByte(a, b *Byte) *Byte {
    return NewByte(a.Value / b.Value)
}


func ModByte(a, b *Byte) *Byte {
    return NewByte(a.Value % b.Value)
}


func (x *Byte) ToBoolean() *Boolean {
    return NewBoolean(x)
}


func (x *Byte) ToInteger() *Integer {
    return NewInteger( int64(x.Value) )
}


func (x *Byte) ToRune() *Rune {
    return NewRune( []byte{x.Value} )
}
