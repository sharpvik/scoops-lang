package queue

import (
    "testing"
    "github.com/sharpvik/scoops/Package/DataTypes/Primitives"
    "github.com/sharpvik/scoops/Package/Shared"
)



func TestPush(t *testing.T) {
    q := New()
    q.Push( primitives.NewByte(1) )
    q.Push( primitives.NewByte(2) )
    q.Push( primitives.NewByte(3) )
    if q.Size() != 3 {
        t.Error("Method Push failed test.")
    }
}


func TestEmpty(t *testing.T) {
    q := New()
    a := q.Empty()
    if !a {
        t.Error("Queue is empty but not reported as such.")
    }
    
    q.Push( primitives.NewByte(1) )
    b := q.Empty()
    if b {
        t.Error("Queue is not empty, but reported a empty.")
    }
}


func TestClear(t *testing.T) {
    q := New()
    q.Push( primitives.NewByte(1) )
    q.Push( primitives.NewByte(2) )
    q.Push( primitives.NewByte(3) )
    q.Clear()
    if !q.Empty() {
        t.Error("Clear function failed to clear the queue.")
    }
}


func TestPop(t *testing.T) {
    q := New()
    cases := []shared.Object{
        primitives.NewByte(1), primitives.NewBoolean( primitives.NewByte(1) ),
        primitives.NewByte(3), primitives.NewBoolean( primitives.NewByte(0) ), 
        primitives.NewRune(  []byte{100}  ),
    }
    for _, c := range cases {
        q.Push(c)
    }
    for _, j := range cases {
        i := q.Pop()
        if i != j {
            t.Error("Function Pop returned incorrect value.")
        }
    }
}


func TestPeek(t *testing.T) {
    q := New()
    cases := []shared.Object{
        primitives.NewByte(1), primitives.NewBoolean( primitives.NewByte(1) ),
        primitives.NewByte(3), primitives.NewBoolean( primitives.NewByte(0) ), 
        primitives.NewRune(  []byte{100}  ),
    }
    for _, c := range cases {
        q.Push(c)
    }
    for _, j := range cases {
        i := q.Peek()
        if i != j {
            t.Error("Function Peek returned incorrect value.")
        }
        q.Pop()
    }
}


func TestClone(t *testing.T) {
    q := New()
    q.Push( primitives.NewByte(1) )
    q.Push( primitives.NewByte(2) )
    q.Push( primitives.NewByte(3) )
    newq := q.Clone().(*Queue)
    for i := 0; i < 3; i++ {
        a := q.Pop().(*primitives.Byte).Value
        b := newq.Pop().(*primitives.Byte).Value
        if a != b {
            t.Error("Function 'Clone' works incorrectly.")
        }
    }
}
