# Package \_string

The implementation of the String data type.
[Wikipedia article.](https://en.wikipedia.org/wiki/String_(computer_science))



## Overview

This package provides the String data type and all related functions.



## Index

```go
type String
    func New(value []*primitives.Rune) *String
    func FromString(from string) *String
    func (s *String) Clone() shared.Object
    func (s *String) Print(w *bufio.Writer)
    func (s *String) Type() string
    func (s *String) Size() uint64
    func Concatenate(a, b *String) *String
```


### func New

```go
func New(value []*primitives.Rune) *String
```

**New** returns pointer to a newly created `String` instance.


### func FromString

```go
func FromString(from string) *String
```

**FromString** accepts Go's embedded `string` and return a pointer to a newly
created `String` instance.


### func Clone

```go
func (s *String) Clone() shared.Object
```

**Clone** returns value of interface type `shared.Object` with underlying type
`*String`. The return value of this method is a new `*String` that is identical
to its parent `*String` in its contents.


### func Print

```go
func (s *String) Print(w *bufio.Writer)
```

**Print** prints our `String` using the specified writer.


### func Type

```go
func (s *String) Type() string
```

**Type** returns `String`'s type as Go's embedded `string` type.


### func Size

```go
func (s *String) Size() uint64
```

**Size** returns the number of `Rune`s that are contained within the `String`.


### func Concatenate

```go
func Concatenate(a, b *String) *String
```

**Concatenate** accepts two `String`s A and B and returns a `String` that is
made up of the input `String`s in their respectable order.
