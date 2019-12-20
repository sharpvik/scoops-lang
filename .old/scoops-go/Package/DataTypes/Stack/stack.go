package stack

import (
    "bufio"
    "github.com/sharpvik/scoops/Package/Shared"
)



type (
    node struct {
        prev    *node
        next    *node
        val     shared.Object
    }

    Stack struct {
        size    uint64
        top     *node
        bottom  *node
    }
)



func newNode(prev, next *node, val shared.Object) *node {
    return &node{prev, next, val}
}


func New() *Stack {
    return &Stack{0, nil, nil}
}


func (s *Stack) Clear() {
    s.size = 0
    s.top = nil
    s.bottom = nil
}


func (s *Stack) Clone() shared.Object {
    news := New()
    if s.size == 0 {
        return news
    }
    cur := s.bottom
    for {
        news.Push( cur.val.Clone() )
        if cur == s.top {
            break
        }
        cur = cur.prev
    }
    return news
}


func (s *Stack) Empty() bool {
    return s.size == 0
}


func (s *Stack) Peek() shared.Object {
    return s.top.val
}


func (s *Stack) Pop() shared.Object {
    if s.size == 0 {
        return nil
    }
    n := s.top.val
    if s.size == 1 {
        s.Clear()
        return n
    }
    s.top = s.top.next
    s.top.prev = s.bottom
    s.bottom.next = s.top
    s.size--
    return n
}


func (s *Stack) Print(w *bufio.Writer) {
    defer w.Flush()
    if s.size == 0 {
        w.WriteString("<-> [ ]")
        return
    }
    cur := s.top
    w.WriteString("<-> [ ")
    for {
        cur.val.Print(w)
        w.WriteByte(' ')
        if cur == s.bottom {
            w.WriteByte(']')
            break
        }
        cur = cur.next
    }
}


func (s *Stack) Push(item shared.Object) {
    if s.size == 0 {
        s.top = newNode(nil, nil, item)
        s.bottom = s.top
        s.top.prev = s.top
        s.top.next = s.top
    } else {
        s.top = newNode(s.bottom, s.top, item)
        s.bottom.next = s.top
        s.top.next.prev = s.top
    }
    s.size++
}


func (s *Stack) Size() uint64 {
    return s.size
}


func (s *Stack) Type() string {
    return "stack"
}
