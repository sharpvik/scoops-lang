package slice

import (
    "testing"
    "github.com/sharpvik/scoops/Package/DataTypes/Primitives"
)



func TestClone(t *testing.T) {
    s := New()
    s.Append( primitives.NewByte(1) )
    s.Append( primitives.NewByte(2) )
    s.Append( primitives.NewByte(3) )
    news := s.Clone().(*Slice)
    for i := 0; i < 3; i++ {
        a := s.Pop(0).(*primitives.Byte).Value
        b := news.Pop(0).(*primitives.Byte).Value
        if a != b {
            t.Error("Function 'Clone' works incorrectly.")
        }
    }
}