package primitives

import (
    "bufio"
    "github.com/sharpvik/scoops/Package/Shared"
    "os"
    "testing"
)


func TestPrint(t *testing.T) {
    wrtr := bufio.NewWriter(os.Stdout)
    e := NewError(shared.OpcodeError, "Unknown opcode found.")
    e.Print(wrtr)
}