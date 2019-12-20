package gofunc

import (
	"bufio"
	"fmt"
	"github.com/sharpvik/scoops/Package/Bytes"
	"github.com/sharpvik/scoops/Package/Shared"
)


type GoFunc struct {
    Name string
    Func func(bytes.Environment)
}


func New( name string, fun func(bytes.Environment) ) *GoFunc {
	return &GoFunc{name, fun}
}


func (f *GoFunc) Clone() shared.Object {
	return New(f.Name, f.Func)
}


func (f *GoFunc) Print(w *bufio.Writer) {
	w.WriteString( fmt.Sprintf("<GoFunc %s at %p>", f.Name, f) )
}


func (f *GoFunc) Type() string {
	return "gofunc"
}
