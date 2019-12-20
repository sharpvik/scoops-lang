package primitives

import (
    "bufio"
    "fmt"
    "github.com/sharpvik/scoops/Package/Shared"
)



type Boolean struct {
    Value bool
}


func NewBoolean(value *Byte) *Boolean {
    return &Boolean{value.Value != byte(0)}
}


func FromBoolean(value bool) *Boolean {
    return &Boolean{value}
}


func (b *Boolean) Clone() shared.Object {
    return FromBoolean(b.Value)
}


func (b *Boolean) Print(w *bufio.Writer) {
    w.WriteString( fmt.Sprintf("%v", b.Value) )
    w.Flush()
}


func (b *Boolean) Type() string {
    return "bln"
}


func NotBoolean(b *Boolean) *Boolean {
    return &Boolean{!b.Value}
}


func AndBoolean(a, b *Boolean) *Boolean {
    return &Boolean{a.Value && b.Value}
}


func OrBoolean(a, b *Boolean) *Boolean {
    return &Boolean{a.Value || b.Value}
}


func XorBoolean(a, b *Boolean) *Boolean {
    return &Boolean{a.Value != b.Value}
}
