package list

import (
    "testing"
    "github.com/sharpvik/scoops/Package/DataTypes/Primitives"
)



func TestAppend(t *testing.T) {
    lst := New()
    if lst.Size() != 0 {
        t.Error("LinkedList construction function 'New' works incorrectly.")
    }
    lst.Append( primitives.NewByte(1) )
    lst.Append( primitives.NewByte(2) )
    lst.Append( primitives.NewByte(3) )
    if lst.Size() != 3 {
        t.Error("Function 'Append' works incorrectly.")
    }
}


func (l *LinkedList) TestGetItemByIndex(t *testing.T) {
    lst := New()
    lst.Append( primitives.NewByte(1) )
    lst.Append( primitives.NewByte(2) )
    lst.Append( primitives.NewByte(3) )
    item, _ := lst.GetItemByIndex(1)
    if item.(*primitives.Byte).Value != 2 {
        t.Error("Function 'GetItemByIndex' works incorrectly.")
    }
}


func (l *LinkedList) TestReassignIndex(t *testing.T) {
    lst := New()
    lst.Append( primitives.NewByte(1) )
    lst.Append( primitives.NewByte(2) )
    lst.Append( primitives.NewByte(3) )
    lst.ReassignIndex(
        1,
        primitives.NewByte(42),
    )
    item, _ := lst.GetItemByIndex(1)
    if item.(*primitives.Byte).Value != 42 {
        t.Error("Function 'ReassignIndex' works incorrectly.")
    }
}


func TestPop(t *testing.T) {
    lst := New()
    lst.Append( primitives.NewByte(1) )
    lst.Append( primitives.NewByte(2) )
    lst.Append( primitives.NewByte(3) )
    lst.Pop(1)
    if lst.Size() != 2 {
        t.Error("Function 'Pop' works incorrectly.")
    }
}
