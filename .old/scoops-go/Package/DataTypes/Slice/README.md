# Package slice

The implementation of the Dynamically Allocated Array data type.
[Wikipedia article.](https://en.wikipedia.org/wiki/Dynamic_array)



## Overview

This package provides the Dynamic Array data type and all related functions.



## Index

```go
type Slice
    func New() *Slice
    func (s *Slice) Clone() shared.Object
    func (s *Slice) Print(w *bufio.Writer)
    func (s *Slice) Type() string
    func (s *Slice) Size() uint64
    func (s *Slice) Append(item shared.Object)
    func (s *Slice) GetItemByIndex(index uint64) shared.Object
    func (s *Slice) Pop(index uint64) shared.Object
```


### func New

```go
func New() *Slice
```

**New** returns pointer to a newly created `Slice` instance.


### func Clone

```go
func (s *Slice) Clone() shared.Object
```

**Clone** returns value of interface type `shared.Object` with underlying type
`*Slice`. The return value of this method is a new `*Slice` that is identical to
its parent `*Slice` in its contents, as every item has also been Cloned.


### func Print

```go
func (s *Slice) Print(w *bufio.Writer)
```

**Print** prints our `Slice` using the specified writer.


### func Type

```go
func (s *Slice) Type() string
```

**Type** returns `Slice`'s type as Go's embedded `string` type.


### func Size

```go
func (s *Slice) Size() uint64
```

**Size** returns the number of elements that are contained within the `Slice`.


### func Append

```go
func (s *Slice) Append(item shared.Object)
```

**Append** adds given element to the end of the `Slice`.


### func GetItemByIndex

```go
func (s *Slice) GetItemByIndex(index uint64) shared.Object
```

**GetItemByIndex** returns item with given index as `shared.Object` interface
type, allowing user to access elements of our `Slice`.


### func Pop

```go
func (s *Slice) Pop(index uint64) shared.Object
```

**Pop** removes element with given index from the `Slice` and returns it as a
`shared.Object`.