package primitives

import (
    "bufio"
    "github.com/sharpvik/scoops/Package/Shared"
)



type Rune struct {
    Value []byte
}


func NewRune(value []byte) *Rune {
    return &Rune{value}
}


func (r *Rune) Clone() shared.Object {
    return NewRune(r.Value)
}


func (r *Rune) Print(w *bufio.Writer) {
    w.Write(r.Value)
    w.Flush()
}


func (r *Rune) Type() string {
    return "rune"
}
