package tuple

import (
    "bufio"
    "github.com/sharpvik/scoops/Package/Shared"
)



type Tuple struct {
	size 	uint64
	value 	[]shared.Object
}



func New(items []shared.Object) *Tuple {
	return &Tuple{uint64( len(items) ), items}
}


func (t *Tuple) Clone() shared.Object {
	var clonedValue []shared.Object
	for _, item := range t.value {
		clonedValue = append( clonedValue, item.Clone() )
	}
	return New(clonedValue)
}


func (t *Tuple) Print(w *bufio.Writer) {
	defer w.Flush()
	if t.size == 0 {
		w.WriteString("[ ]")
		return
	}
	w.WriteString("[ ")
	for _, item := range t.value {
		item.Print(w)
		w.WriteByte(' ')
	}
	w.WriteByte(']')
}


func (t *Tuple) Type() string {
	return "tuple"
}


func (t *Tuple) Size() uint64 {
	return t.size
}