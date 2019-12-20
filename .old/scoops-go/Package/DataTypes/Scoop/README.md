# Package scoop

This package contains the `Scoop` type and methods related to it. `Scoop`s are
used to implement function-like behaviour like in other well-known programming
languages. Nevertheless, the concept of a `Scoop` is not entirely synonymous
with the idea of a function in high-level code.



## Overview

In the *Scoops Programming Language*, `Scoop` is just another kind of object --
it can be stored in a variable, `Push`ed and `Pop`ed off the `data` stack, etc.
`Scoop type` fully implements the `shared.Object interface` and the
`shared.Collection interface` types.

*De facto*, `Scoop` is just a **container** for the code of type
`[]*shared.Instruction`. When `Scoop` is called, the code it contains is put
into the `code` field of the newly created `Environment`. See 
[types.go](../../Bytes/types.go) and [Bytes/README.md](../../Bytes/README.md)
to learn more about the way *Scoops Programming Language* works with execution
scopes.



## Index

```go
type Scoop
    func New(name string, code []*shared.Instruction) *Scoop
    func (s *Scoop) Clone() shared.Object
    func (s *Scoop) Print(w *bufio.Writer)
    func (s *Scoop) Type() string
```


### func New

```go
func New() *Scoop
```

**New** returns pointer to a newly created `Scoop` instance.


### func Clone

```go
func (s *Scoop) Clone() shared.Object
```

**Clone** returns an object of type shared.Object (with underlying type
`*Scoop`) that is identical, in terms of contents, to the `Scoop` this function
was called on.


### func Print

```go
func (s *Scoop) Print(w *bufio.Writer)
```

**Print** prints out some explanatory data about the `Scoop` using the specified
writer.


### func Type

```go
func (s *Scoop) Type() string
```

**Type** returns `Scoop`'s type as Go's embedded `string` type.