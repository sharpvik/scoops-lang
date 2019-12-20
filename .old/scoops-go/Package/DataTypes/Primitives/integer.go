package primitives

import (
    "bufio"
    "fmt"
    "github.com/sharpvik/scoops/Package/Shared"
)


type Integer struct {
    Value int64
}


func NewInteger(value int64) *Integer {
    return &Integer{value}
}


func (i *Integer) Clone() shared.Object {
    return NewInteger(i.Value)
}


func (i *Integer) Print(w *bufio.Writer) {
    w.WriteString( fmt.Sprintf("%d", i.Value) )
    w.Flush()
}


func (i *Integer) Type() string {
    return "int"
}


func AddInteger(a, b *Integer) *Integer {
    return NewInteger(a.Value + b.Value)
}


func SubInteger(a, b *Integer) *Integer {
    return NewInteger(a.Value - b.Value)
}


func MulInteger(a, b *Integer) *Integer {
    return NewInteger(a.Value * b.Value)
}


func DivInteger(a, b *Integer) *Float {
    return NewFloat( float64(a.Value) / float64(b.Value) )
}


func ModInteger(a, b *Integer) *Integer {
    return NewInteger(a.Value % b.Value)
}


func (i *Integer) ToFloat() *Float {
    return NewFloat( float64(i.Value) )
}