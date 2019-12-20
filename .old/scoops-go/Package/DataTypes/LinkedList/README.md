# Package tuple

This package contains the `LinkedList` type and methods related to it. 
`LinkedList` is a collection of elements -- just like a `Slice` -- with one
major difference -- its elements are not stored in memory *continuously*.



## Overview

`LinkedList` is just a collection of `node`s that hold three things each:

1. Pointer to the previous `node`;
2. Pointer to the next `node`;
3. The element itself (of type `shared.Object`)

`LinkedList` type itself only holds its size and a pointer to the first `node`.
This data structure is very dynamic as it grows only when new elements are added
to it and its growth is cheaper than that of a dynamic array since it doesn't
involve asking computer to allocate more memory or copying elements into a newly
allocated space. However, the fact that it's discontinuous means that getting
the value of an element under a certain index is no longer O(1) -- it's O(n).



## Index

```go
type LinkedList
    func New() *LinkedList
    func (l *LinkedList) Clone() shared.Object
    func (l *LinkedList) Print(w *bufio.Writer)
    func (l *LinkedList) Type() string
    func (l *LinkedList) Size() uint64
    func (l *LinkedList) Append(item shared.Object)
    func (l *LinkedList) Clear()
    func (l *LinkedList) Empty() bool
    func (l *LinkedList) GetItemByIndex(index uint64) 
                                       (shared.Object, *primitives.Error)
    func (l *LinkedList) ReassignIndex(index uint64, item shared.Object)
                                      (*primitives.Error)
    func (l *LinkedList) Pop(index uint64) (shared.Object, *primitives.Error)
```


### func New

```go
func New() *LinkedList
```

**New** returns pointer to a newly created `LinkedList` instance.


### func Clone

```go
func (l *LinkedList) Clone() shared.Object
```

**Clone** returns an object of type shared.Object (with underlying type
`*LinkedList`) that is identical, in terms of contents, to the `LinkedList` this 
function was called on.


### func Print

```go
func (l *LinkedList) Print(w *bufio.Writer)
```

**Print** prints out `LinkedList`s contents using a specified writer.


### func Type

```go
func (l *LinkedList) Type() string
```

**Type** returns `LinkedList`'s type as Go's embedded `string` type.


### func Size

```go
func (l *LinkedList) Size() uint64
```

**Size** returns the number of elements in the `LinkedList`.


### func Append

```go
func (l *LinkedList) Append(item shared.Object)
```

**Append** adds given element to the end of the `LinkedList`.


### func Clear

```go
func (l *LinkedList) Clear()
```

**Clear** removes all elements from our `LinkedList`. Given the nature of this
data structure, the **Empty** function is *much* more efficient than removing 
every single element using the **Pop** method.


### func Empty

```go
func (l *LinkedList) Empty() bool
```

**Empty** returns `true` if `LinkedList` is empty, otherwise returns `false`.


### func GetItemByIndex

```go
func (l *LinkedList) GetItemByIndex(index uint64) 
                                   (shared.Object, *primitives.Error)
```

**GetItemByIndex** returns item with given index as `shared.Object` interface
type, allowing user to access elements of our `LinkedList`. Non-`nil` second 
return value of type `*primitives.Error` signifies that an error occurred while
trying to get the item.


### func ReassignIndex

```go
func (l *LinkedList) ReassignIndex(index uint64, item shared.Object)
                                  (*primitives.Error)
```

**ReassignIndex** changes the value of a node with given index to `item`. If
any error occurs, it returns some `*primitives.Error`.


### func Pop

```go
func (l *LinkedList) Pop(index uint64) shared.Object
```

**Pop** removes element with given index from the `LinkedList` and returns it
as a `shared.Object`.