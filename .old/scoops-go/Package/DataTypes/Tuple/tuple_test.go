package tuple

import (
	"github.com/sharpvik/scoops/Package/Shared"
	"bufio"
	"os"
	"testing"
    "github.com/sharpvik/scoops/Package/DataTypes/Primitives"
)



func TestClone(t *testing.T) {
	tup := New([]shared.Object{
		primitives.NewByte(1),
		primitives.NewByte(2),
		primitives.NewByte(3),
	})
	ntup := tup.Clone().(*Tuple)
	w := bufio.NewWriter(os.Stdout)
	tup.Print(w)
	w.WriteByte('\n')
	ntup.Print(w)
	w.WriteByte('\n')
	w.Flush()
}