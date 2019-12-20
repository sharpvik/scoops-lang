# Package tuple

This package contains the `Tuple` type and methods related to it. `Tuple`s are
immutable `Slice`s -- they can only be created and cloned. The idea behind a
`Tuple` is to provide a constant (hence *hashable*) data type to programmers.



## Overview

As pointed out by some smart man from YouTube, hash-tables are very handy for a
variety of reasons, yet they have a severe limitation -- they can only use
hashable data types as keys. Basically, it means that you can practically only
ever use `string`s and `integer` types as keys. That is unless, in your
programming language, you have a collection that is immutable. In Python, for
example, you have `Tuple` type given to you by default which allows you to use
collection as a key. I think we should adopt this great idea into our family.



## Index

```go
type Tuple
    func New(items []shared.Object) *Tuple
    func (t *Tuple) Clone() shared.Object
    func (t *Tuple) Print(w *bufio.Writer)
    func (t *Tuple) Type() string
    func (t *Tuple) Size() uint64
```


### func New

```go
func New(items []shared.Object) *Tuple
```

**New** returns pointer to a newly created `Tuple` instance.


### func Clone

```go
func (t *Tuple) Clone() shared.Object
```

**Clone** returns an object of type shared.Object (with underlying type
`*Tuple`) that is identical, in terms of contents, to the `Tuple` this function
was called on.


### func Print

```go
func (t *Tuple) Print(w *bufio.Writer)
```

**Print** prints out `Tuple`s contents using a specified writer.


### func Type

```go
func (t *Tuple) Type() string
```

**Type** returns `Tuple`'s type as Go's embedded `string` type.


### func Size

```go
func (t *Tuple) Size() uint64
```

**Size** returns the number of elements in the `Tuple`.