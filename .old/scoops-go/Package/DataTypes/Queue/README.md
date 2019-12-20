# Package queue

The implementation of the Queue abstract data type.
[Wikipedia article.](https://en.wikipedia.org/wiki/Queue_(abstract_data_type))



## Overview

This package provides the Queue data type and functions that are related to it.



## Index

```go
type Queue
    func New() *Queue
    func (q *Queue) Clear()
    func (q *Queue) Clone() shared.Object
    func (q *Queue) Empty() bool
    func (q *Queue) Peek() shared.Object
    func (q *Queue) Pop() shared.Object
    func (q *Queue) Print(w *bufio.Writer)
    func (q *Queue) Push(item shared.Object)
    func (q *Queue) Size() uint64
    func (q *Queue) Type() string
```


### func New

```go
func New() *Queue
```

**New** returns pointer to a newly initialised Queue.


### func Clear

```go
func (q *Queue) Clear()
```

**Clear** simply empties the queue. It is *much* more efficient to use **Clear**
than popping every element off by hand using the **Pop** function.


### func Clone

```go
func (q *Queue) Clone() shared.Object
```

**Clone** returns an object of type shared.Object (with underlying type
`*Queue`) that is identical, in terms of contents, to the `Queue` this function
was called on.


### func Empty

```go
func (q *Queue) Empty() bool
```

**Empty** returns `true` if queue is empty and `false` if it isn't.


### func Peek

```go
func (q *Queue) Peek() shared.Object
```

**Peek** returns the top item of the queue. Top element is said to be the
element that would be popped if **Pop** function was called.


### func Pop

```go
func (q *Queue) Pop() shared.Object
```

**Pop** pops off and returns the top item of the queue.


### func Print

```go
func (q *Queue) Print(w *bufio.Writer)
```

**Print** prints out contents of the `Queue` using the specified writer.


### func Push

```go
func (q *Queue) Push(item shared.Object)
```

**Push** inserts given item into the queue.


### func Size

```go
func (q *Queue) Size() uint64
```

**Size** returns the number of elements in the queue.


### func Type

```go
func (q *Queue) Type() string
```

**Type** returns `Queue`'s type as Go's embedded `string` type.



## Examples

```go
package main

import (
    "fmt"
    "github.com/sharpvik/scoops/Package/DataTypes/Queue"
)

func main() {
    q := Queue.New()    // init new Queue and store pointer to it in 'q'
    q.Push( primitives.NewByte(1) ) // push first element into the queue
    q.Print()           // print queue to stdout
    one := q.Peek()     // store the top item of the queue in 'one'
    two := q.Pop()      // pop the top item off and store it in 'two'
    q.Print()           // print now empty queue to stdout
    b := (one == two)   // 'b' is true
    fmt.Println(b)      // print value of 'b' to stdout
}
```