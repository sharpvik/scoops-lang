package primitives

import (
    "bufio"
    "github.com/sharpvik/scoops/Package/Shared"
)



type Nil struct {
    value byte
}


func NewNil() *Nil {
    return &Nil{0}
}


func (n *Nil) Clone() shared.Object {
    return n
}


func (n *Nil) Print(w *bufio.Writer) {
    w.WriteString("nil")
    w.Flush()
}


func (n *Nil) Type() string {
    return "nil"
}
