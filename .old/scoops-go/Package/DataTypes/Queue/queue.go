package queue

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

    Queue struct {
        size    uint64
        head    *node
    }
)



func newNode(prev, next *node, val shared.Object) *node {
    return &node{prev, next, val}
}


func New() *Queue {
    return &Queue{0, nil}
}


func (q *Queue) Clear() {
    q.head = nil
    q.size = 0
}


func (q *Queue) Clone() shared.Object {
    newq := New()
    if q.size == 0 {
        return newq
    }
    cur := q.head
    for {
        newq.Push( cur.val.Clone() )
        cur = cur.next
        if cur == q.head {
            break
        }
    }
    return newq
}


func (q *Queue) Empty() bool {
    return q.size == 0
}


func (q *Queue) Peek() shared.Object {
    if q.size == 0 {
        return nil
    }
    return q.head.val
}


func (q *Queue) Pop() shared.Object {
    var n shared.Object
    if q.size == 0 {
        return nil
    } else if q.size == 1 {
        n = q.head.val
        q.head = nil
    } else {
        n = q.head.val
        q.head.prev.next = q.head.next
        q.head.next.prev = q.head.prev
        q.head = q.head.next
    }
    q.size--
    return n
}


func (q *Queue) Print(w *bufio.Writer) {
    defer w.Flush()
    if q.size == 0 {
        w.WriteString("<- [ ] <-")
        return
    }
    cur := q.head
    w.WriteString("<- [ ")
    for {
        cur.val.Print(w)
        w.WriteString(" ")
        cur = cur.next
        if cur == q.head {
            w.WriteString("] <-")
            break
        }
    }
}


func (q *Queue) Push(item shared.Object) {
    n := newNode(nil, nil, item)
    if q.size == 0 {
        n.prev = n
        n.next = n
        q.head = n
    } else {
        n.prev = q.head.prev
        n.next = q.head
        q.head.prev.next = n
        q.head.prev = n
    }
    q.size++
}


func (q *Queue) Size() uint64 {
    return q.size
}


func (q *Queue) Type() string {
    return "queue"
}
