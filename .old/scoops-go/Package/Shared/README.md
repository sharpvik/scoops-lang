# Package shared

This package contains all sorts of things like types, constants, etc. that need
to be imported into multiple files.



## Overview

[errors.go]     -- error tags
[opcodes.go]    -- opcodes used by the Scoops Virtual Machine
[types.go]      -- interfaces and structs 

[errors.go]:    ./errors.go
[opcodes.go]:   ./opcodes.go
[types.go]:     ./types.go



## Error tags

Error tags are used to track and *classify* errors (sort of like error codes but
non-numeric). There isn't much to say about error tags to be honest, so we'll
move on to something else.



## Opcodes

Opcodes are simply **constants** that hold exactly **one byte**.
*Scoops Virtual Machine* uses these constants to identify opcode for each
instruction it receives and thus, decide what to do.



## Types

Some types and interfaces turned out to me so important that they are needed by
a few different modules, thus they are declared here, in the *Shared* folder.

There are two `interface` types -- `Object` and `Collection` -- that are 
required to provide polymorphism. Type `Object` is precisely the type of objects
stored on the `data stack`. See [Bytes/README.md](../Bytes/README.md) to learn
more about the `Environment type` and `data stack` that it uses.

```go
type Object interface {
    Clone() Object
    Print(w *bufio.Writer)
    Type() string
}

type Collection interface {
    Object
    Size() uint64
}
```

Every `struct` that wants to be able to be `Push`ed onto the `data stack` has to
implement these three basic methods: `Clone`, `Print`, and `Type`. Every
`Collection` implements every method stated in the `Object` interface plus the
`Size` method which is reasonable; it also means that any struct that implements
`Collection` interface, also satisfies the `Object` interface by definition and,
therefore, can also be `Push`ed onto the `stack`.

Function `Print` has a special place in my heart as it allows us to work with
Strings in a way that is better that Python's in terms of memory efficiency and
is also simpler than Go's as it is *conversion free*.

UTF-8 made it sort of impossible to represent strings as byte arrays where every
single byte is a character. Now, characters have *variable length*, which makes 
it harder for us to work with string indexing. Python (as far as I know) 
represents characters as 16-bit-long integers -- which isn't really memory 
efficient -- in order to keep its ability to index characters instead of bytes.
(Go's strings are indexed based on bytes, not characters.)

Go, on the other hand, introduces type rune to represent UTF-8 characters while
it keeps its string type defined as an immutable byte array. This means that 
iterating through a string becomes complicated as programmers have to always 
worry about the 'aggregate state' of their string -- is it of `string` type?; is 
it an array of runes (`[]rune`)?; or maybe it is an array of bytes (`[]byte`)..? 

In each case, iteration and indexing differ. Moreover, if I'm not mistaken, type
rune also follows Python's memory inefficient approach, but they use 32-bit-long
integers... which means that rune array that represents a text comprised of 
ASCII characters only, takes up four times as much space as it actually needs.

As a way to deal with the problems stated above, *Scoops* uses its custom `Rune`
type based on Go's `byte slice`. That way we only store as many bytes as 
necessary, the bytes themselves are always there, ready to be used without any 
conversions, yet `String` is defined as a slice of `Rune`s, so indexing is based
on characters like in Python (which is handy in my opinion). There will be a way
to convert `String` to `Slice` of `Byte`s if you'd like though -- you'll then be
able to iterate over `Byte`s, not `Rune`s. Mentioned conversion is also pretty
cheap as the bytes are *already there* and we just need to collect them. 
However, if you want to print your `String` to `stdout`, file or otherwise,
**no conversion needed**. Use `String`'s `Print` method that simply iterates
over its `Rune`s and prints bytes of each `Rune` to the specified `io.Writer`.