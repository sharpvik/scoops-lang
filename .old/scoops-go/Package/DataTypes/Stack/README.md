# Package stack

The implementation of the Stack abstract data type.
[Wikipedia article.](https://en.wikipedia.org/wiki/Stack_(abstract_data_type))



## Overview

This package provides the Stack data type and functions that are related to it.



## Index

```go
type Stack
    func New() *Stack
    func (s *Stack) Clear()
    func (s *Stack) Clone() shared.Object
    func (s *Stack) Empty() bool
    func (s *Stack) Peek() shared.Object
    func (s *Stack) Pop() shared.Object
    func (s *Stack) Print(w *bufio.Writer)
    func (s *Stack) Push(item shared.Object)
    func New() *Stack
    func (s *Stack) Type() string
```


### func New

```go
func New() *Stack
```

**New** returns pointer to a newly created `Stack` instance.


### func Clear

```go
func (s *Stack) Clear()
```

**Clear** simply empties the stack. It is *much* more efficient to use **Clear**
than popping every element off by hand using the **Pop** function.


### func Clone

```go
func (s *Stack) Clone() shared.Object
```

**Clone** returns an object of type shared.Object (with underlying type
`*Stack`) that is identical, in terms of contents, to the `Stack` this function
was called on.


### func Empty

```go
func (s *Stack) Empty() bool
```

**Empty** returns `true` if stack is empty and `false` if it isn't.


### func Peek

```go
func (s *Stack) Peek() shared.Object
```

**Peek** returns the top item of the stack. Top element is said to be the
element that would be popped if **Pop** function was called.


### func Pop

```go
func (s *Stack) Pop() shared.Object
```

**Pop** pops off and returns the top item of the stack.


### func Print

```go
func (s *Stack) Print(w *bufio.Writer)
```

**Print** prints out contents of the `Stack` using the specified writer.


### func Push

```go
func (s *Stack) Push(item shared.Object)
```

**Push** inserts given item into the stack.


### func Size

```go
func (s *Stack) Size() uint64
```

**Size** returns the number of elements in the stack.


### func Type

```go
func (s *Stack) Type() string
```

**Type** returns `Stack`'s type as Go's embedded `string` type.



## Examples

```go
package main

import (
    "fmt"
    "github.com/sharpvik/scoops/Package/DataTypes/Stack"
)

func main() {
    q := Stack.New()    // init new Stack and store pointer to it in 'q'
    q.Push(1)           // push first element into the stack
    q.Print()           // print stack to stdout
    one := q.Peek()     // store the top item of the stack in 'one'
    two := q.Pop()      // pop the top item off and store it in 'two'
    q.Print()           // print now empty stack to stdout
    b := (one == two)   // 'b' is true
    fmt.Println(b)      // print value of 'b' to stdout
}
```