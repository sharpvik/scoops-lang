# Package primitives



## Overview

This package contains files in which one can find structs that represent the
basic data types (primitives) of the Scoops language and all related methods.
It is important to note that all types declared in this package implement the
`Object` interface from `scoops/Package/Shared`.



## Index

```go
type Boolean
    func NewBoolean(value *Byte) *Boolean
    func FromBoolean(value bool) *Boolean
    func (b *Boolean) Clone() shared.Object
    func (b *Boolean) Print()
    func (b *Boolean) Type() string
    func NotBoolean(b *Boolean) *Boolean
    func AndBoolean(a, b *Boolean) *Boolean
    func OrBoolean(a, b *Boolean) *Boolean
    func XorBoolean(a, b *Boolean) *Boolean
type Byte
    func NewByte(value byte) *Byte
    func (b *Byte) Clone() shared.Object
    func (b *Byte) Print()
    func (b *Byte) Type() string
    func AddByte(a, b *Byte) *Byte
    func SubByte(a, b *Byte) *Byte
    func MulByte(a, b *Byte) *Byte
    func DivByte(a, b *Byte) *Byte
    func (x *Byte) ToBoolean() *Boolean
    func (x *Byte) ToInteger() *Integer
    func (x *Byte) ToRune() *Rune
type Error
    func NewError(tag, message string) *Error
    func FromError(tag string, message error) *Error
    func (e *Error) Clone() *Error
    func (e *Error) Print(w *bufio.Writer)
    func NewError(tag, message string) *Error
func FromError(tag string, message error) *Error
func (e *Error) ToGoString() string
func (e *Error) ToGoError() error
    func (e *Error) ToGoError() error
    func (e *Error) Type() string
type Float
    func NewFloat(value float64) *Float
    func (f *Float) Clone() shared.Object
    func (f *Float) Print(w *bufio.Writer)
    func (f *Float) Type() string
    func AddFloat(a, b *Float) *Float
    func SubFloat(a, b *Float) *Float
    func MulFloat(a, b *Float) *Float
    func DivFloat(a, b *Float) *Float
type Integer
    func NewInteger(value int64) *Integer
    func (i *Integer) Clone() shared.Object
    func (i *Integer) Print()
    func (i *Integer) Type() string
    func AddInteger(a, b *Integer) *Integer
    func SubInteger(a, b *Integer) *Integer
    func MulInteger(a, b *Integer) *Integer
    func DivInteger(a, b *Integer) *Float
type Nil
    func NewNil() *Nil
    func (n *Nil) Clone() shared.Object
    func (n *Nil) Print()
    func (n *Nil) Type() string
type Rune
    func NewRune(value []byte) *Rune
    func (r *Rune) Clone() shared.Object
    func (r *Rune) Print()
    func (r *Rune) Type() string
```



## Functions that belong to shared.Object interface

### func Clone

```go
func (b *<primitive>) Clone() shared.Object
```

**Clone** returns value of interface type `shared.Object` that can be casted
back into its primitive type using a type assertion like so:

```go
/* The var <varname> <data type> notation is used here to be explicit about the
   data types of those variables. This is just a demo. */
var a *Boolean = FromBoolean(true) // creating an instance of Boolean
var c shared.Object = a.Clone() // Clone function returns shared.Object
var b *Boolean = c.(*Boolean) // casting shared.Object back into *Boolean
```

Of course, the same thing can be done in a much shorter way:

```go
var a *Boolean = FromBoolean(true) // creating an instance of Boolean
var b *Boolean = a.Clone().(*Boolean) // cloning a into b
```


### func Print

```go
func (b *<primitive>) Print()
```

**Print** function defines the default way of displaying (printing) a primitive
to the stdout.


### func Type

```go
func (b *<primitive>) Type() string
```

**Type** returns primitive's type as Go's embedded `string` type. The `GET_TYPE`
opcode uses `string.FromString` function from the
`scoops/Package/DataTypes/String` package to convert `string` to
`_string.String` type and push it onto the data stack.



## Type Boolean

### func NewBoolean

```go
func NewBoolean(value *Byte) *Boolean
```

**NewBoolean** returns pointer to a newly created `Boolean` instance.


### func FromBoolean

```go
func FromBoolean(value bool) *Boolean
```

**FromBoolean** returns pointer to a newly created `Boolean` instance. It is
used when we want to create a `*Boolean` using Go's embedded `bool` data type as
an argument to the function.


### func NotBoolean, AndBoolean, OrBoolean, XorBoolean

```go
func NotBoolean(a *Boolean) *Boolean
func AndBoolean(a, b *Boolean) *Boolean
func OrBoolean(a, b *Boolean) *Boolean
func XorBoolean(a, b *Boolean) *Boolean
```

All functions listed above implement basic boolean functions that are essential
for any programming language.
**NotBoolean** reverses the value of the incoming
`*Boolean`.
**AndBoolean** accepts two `*Boolean` values and returns a `*Boolean` with a
`true` value if both incoming booleans were `true`.
**OrBoolean** also accepts two `*Boolean`s and returns `true` if at least one of
those was `true`.
**XorBoolean** accepts two `*Boolean`s -- `a` and `b` -- and returns `true` if
`a` equals `not b`, or in other words, if `a` not the same as `b`.



## Type Byte

### func NewByte

```go
func NewByte(value byte) *Byte
```

**NewByte** returns pointer to a newly created `Byte` instance.


### func AddByte, SubByte, MulByte, DivByte

```go
func AddByte(a, b *Byte) *Byte
func SubByte(a, b *Byte) *Byte
func MulByte(a, b *Byte) *Byte
func DivByte(a, b *Byte) *Byte
```

Functions listed above implement `*Byte`-based arithmetic.
In case you want to add two `*Byte`s -- use **AddByte**;
In case you want to subtract two `*Byte`s -- use **SubByte**;
In case you want to multiply two `*Byte`s -- use **MulByte**;
In case you want to divide two `*Byte`s -- use **DivByte**.


### func ToBoolean, ToInteger, ToRune

```go
func (x *Byte) ToBoolean() *Boolean
func (x *Byte) ToInteger() *Integer
func (x *Byte) ToRune() *Rune
```

Functions listed above can convert `*Byte` into other Scoops data types.
**ToBoolean** converts any non-zero `*Byte` into `true`.
**ToInteger** is pretty simple -- just makes it a normal `*Integer`.
**ToRune** converts `*Byte` into a `*Rune` that represents some ASCII character.



## Type Error

### func NewError

```go
func NewError(tag, message string) *Error
```

**NewError** returns pointer to a newly created `Error` instance.


### func FromError

```go
func FromError(tag string, message error) *Error
```

**FromError** is a function that accepts error message as Go's embedded `error`
type. That is its only difference from **NewError**.


### func ToGoString

```go
func (e *Error) ToGoString() string
```

**ToGoString** combines fields of `Error` to return a formatted `string` that
may be used to display that `Error`.


### func ToGoError

```go
func (e *Error) ToGoError() error
```

**ToGoError** simply converts the output of **ToGoString** into Go's embedded
`error` data type and returns the result of that conversion.



## Type Float

### func NewFloat

```go
func NewFloat(value float64) *Float
```

**NewFloat** returns pointer to a newly created `Float` instance.


### func AddFloat, SubFloat, MulFloat, DivFloat

```go
func AddFloat(a, b *Float) *Float
func SubFloat(a, b *Float) *Float
func MulFloat(a, b *Float) *Float
func DivFloat(a, b *Float) *Float
```

Functions above are used to add, subtract, multiply or divide two `Float`s. I
think this is self-explanatory.



## Type Integer

### func NewInteger

```go
func NewInteger(value int64) *Integer
```

**NewInteger** returns pointer to a newly created `Integer` instance.


### func AddInteger, SubInteger, MulInteger

```go
func AddInteger(a, b *Integer) *Integer
func SubInteger(a, b *Integer) *Integer
func MulInteger(a, b *Integer) *Integer
func DivInteger(a, b *Integer) *Float
```

Functions listed above implement `*Integer`-based arithmetic.
In case you want to add two `*Integer`s -- use **AddInteger**;
In case you want to subtract two `*Integer`s -- use **SubInteger**;
In case you want to multiply two `*Integer`s -- use **MulInteger**;
**DivInteger** returns a `*Float` as a result of dividing two `*Integer`s;



## Type Nil

### func NewNil

```go
func NewNil() *Nil
```

**NewNil** returns pointer to a newly created `Nil` instance. `Nil` is my
favourite -- I don't need to write a lot of things about it as it doesn't have
its own functions. `Nil` only implements the functions that are required to
satisfy the `shared.Object` interface.

Now, you may ask yourself: "What's the point of `Nil` data type?"
Well, it is an interesting question, my dear. It is a symbolic value used to
signify that something is nothing. I know, it's a vague explanation, but keep
listening. When there is a scoop that has no return value, it returns `Nil` to
its parent scope. If you create a variable that this scoop's return value must
be assigned to, like that

```scoops
n := IReturnNil()
```

then the variable `n` will hold the `Nil` value. If you want to declare a
variable that doesn't hold any specific value, you may assign it to `Nil`.



## Type Rune

### func NewRune

```go
func NewRune(value []byte) *Rune
```

**NewRune** returns pointer to a newly created `Rune` instance. `Rune` only
implements the functions that are required to satisfy the `shared.Object`
interface. Therefore, I can just mention that `Rune` is just a type that holds
a few `byte`s (from 1 to 4 to be precise) and represent a printable UTF-8
character. Scoops' `Rune` type is a nice compromise between Go's `byte`s,
`string`s, and `rune`s. `_string.String` made of `Rune`s is very flexible:

1. It can be indexed like a Python string;
2. It still holds the `byte`s needed to print it to a file;
3. All of the above -- conversion free!



## Examples

You can find examples in [this folder].

[this folder]: ../../InputFiles/scpa