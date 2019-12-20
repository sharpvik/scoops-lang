package scoop

import (
	"bufio"
	"fmt"
	"github.com/sharpvik/scoops/Package/Shared"
)



type Scoop struct {
	Name string
	Size uint64
	Code []*shared.Instruction
}


func New(name string, code []*shared.Instruction) *Scoop {
	return &Scoop{name, uint64( len(code) ), code}
}


func (s *Scoop) Clone() shared.Object {
	return New(s.Name, s.Code)
}


func (s *Scoop) Print(w *bufio.Writer) {
	w.WriteString(
		fmt.Sprintf("<Scoop %s at %p with size of %d>", s.Name, s, s.Size),
	)
}


func (s *Scoop) Type() string {
	return "scoop"
}
