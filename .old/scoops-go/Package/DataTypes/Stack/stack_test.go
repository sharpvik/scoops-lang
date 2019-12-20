package stack

import (
    "testing"
    "github.com/sharpvik/scoops/Package/DataTypes/Primitives"
)



func TestPush(t *testing.T) {
    s := New()
    s.Push( primitives.NewByte(1) )
    s.Push( primitives.NewByte(2) )
    s.Push( primitives.NewByte(3) )
    if s.size != 3 {
        t.Error("Function 'Push' works incorrectly.")
    }
}


func TestEmpty(t *testing.T) {
    s := New()
    if !s.Empty() {
        t.Error("Function 'Empty' works incorrectly.")
    }
    s.Push( primitives.NewByte(1) )
    if s.Empty() {
        t.Error("Function 'Empty' works incorrectly.")
    }
}


func TestClear(t *testing.T) {
    s := New()
    s.Push( primitives.NewByte(1) )
    s.Push( primitives.NewByte(2) )
    s.Push( primitives.NewByte(3) )
    s.Clear()
    if !s.Empty() {
        t.Error("Funciton 'Clear' works incorrectly.")
    }
}


func TestPop(t *testing.T) {
    s := New()
    s.Push( primitives.NewByte(1) )
    s.Push( primitives.NewByte(2) )
    s.Push( primitives.NewByte(3) )
    v := s.Pop()
    if s.Size() != 2 || v.(*primitives.Byte).Value != byte(3) {
        t.Error("Function 'Pop' works incorrectly.")
    }
    s.Pop()
    v = s.Pop()
    if !s.Empty() || v.(*primitives.Byte).Value != byte(1) {
        t.Error("Function 'Pop' works incorrectly.")
    }
}


func TestClone(t *testing.T) {
    s := New()
    s.Push( primitives.NewByte(1) )
    s.Push( primitives.NewByte(2) )
    s.Push( primitives.NewByte(3) )
    news := s.Clone().(*Stack)
    for i := 0; i < 3; i++ {
        a := s.Pop().(*primitives.Byte).Value
        b := news.Pop().(*primitives.Byte).Value
        if a != b {
            t.Error("Function 'Clone' works incorrectly.")
        }
    }
}