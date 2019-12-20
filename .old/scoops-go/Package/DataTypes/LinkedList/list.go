package list

import (
    "bufio"
    "github.com/sharpvik/scoops/Package/DataTypes/Primitives"
    "github.com/sharpvik/scoops/Package/Shared"
)



type (
    node struct {
        index   uint64
        prev    *node
        next    *node
        val     shared.Object
    }

    LinkedList struct {
        size    uint64
        begin   *node
    }
)


func newNode(index uint64, prev, next *node, val shared.Object) *node {
    return &node{index, prev, next, val}
}


func New() *LinkedList {
    return &LinkedList{0, nil}
}


func (l *LinkedList) Clone() shared.Object {
    newl := New()
    if l.size == 0 {
        return newl
    }
    cur := l.begin
    for {
        newl.Append( cur.val.Clone() )
        if cur == l.begin.prev {
            break
        }
        cur = cur.next
    }
    return newl
}


func (l *LinkedList) Print(w *bufio.Writer) {
    defer w.Flush()
    if l.size == 0 {
        w.WriteString("[ ]")
        return
    }
    w.WriteString("[ ")
    cur := l.begin
    for {
        obj := cur.val
        obj.Print(w)
        w.WriteByte(' ')
        if cur == l.begin.prev {
            break
        }
        cur = cur.next
    }
    w.WriteByte(']')
}


func (l *LinkedList) Type() string {
    return "list"
}


func (l *LinkedList) Size() uint64 {
    return l.size
}


func (l *LinkedList) Append(item shared.Object) {
    if l.size == 0 {
        n := newNode(0, nil, nil, item)
        n.prev, n.next = n, n
        l.begin = n
    } else {
        n := newNode(l.size, l.begin.prev, l.begin, item)
        l.begin.prev.next = n
        l.begin.prev = n
    }
    l.size++
}


func (l *LinkedList) Clear() {
    l.size = 0
    l.begin = nil
}


func (l *LinkedList) Empty() bool {
    return l.size == 0
}


func (l *LinkedList) GetItemByIndex(index uint64) (
        shared.Object,
        *primitives.Error,
    ) {
    if index >= l.size {
        return nil, primitives.NewError(
            shared.IndexError,
            "Index given in call to 'GetItemByIndex' is out of range.",
        )
    }
    cur := l.begin
    for cur.index != index {
        cur = cur.next
    }
    return cur.val, nil
}


func (l *LinkedList) ReassignIndex(index uint64, item shared.Object) (
        *primitives.Error,
    ) {
    if index >= l.size {
        return primitives.NewError(
            shared.IndexError,
            "Index given in call to 'ReassignIndex' is out of range.",
        )
    }
    cur := l.begin
    for cur.index != index {
        cur = cur.next
    }
    cur.val = item
    return nil
}


func (l *LinkedList) Pop(index uint64) (shared.Object, *primitives.Error) {
    len := l.begin.prev.index + 1
    if index >= len {
        return nil, primitives.NewError(
            shared.IndexError,
            "Index given in call to 'Pop' is out of range.",
        )
    }
    if l.size == 1 {
        obj := l.begin.val
        l.Clear()
        l.size--
        return obj, nil
    }
    cur := l.begin
    for cur.index != index {
        cur = cur.next
    }
    cur.prev.next = cur.next
    cur.next.prev = cur.prev
    if index == 0 {
        l.begin = cur.next
    }
    l.size--
    return cur.val, nil
}
