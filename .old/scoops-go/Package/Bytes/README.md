# Package bytes

This package provides structs, methods, and functions that deal with bytecode
execution.



## Overview

The two most important files in this folder are [types.go](./types.go) and
[interpreter.go](./interpreter.go), although, it would be great if you were to
look at every single file.



## Types

In order to make our interpreter smart and flexible, to support execution 
scopes, we have to have some data structures in place. Interpreter *itself* is a
data structure with a couple of methods.

Now, let's look at the two fundamental types separately:

* Interpreter
* Environment


### Interpreter type

```go
type Interpreter struct {
        running bool
        err     *primitives.Error
        global  *Environment
        scope   *Environment    // current execution scope
        thenil  *primitives.Nil // universal nil value
        stdout  *bufio.Writer
        writer  *bufio.Writer   // current writer (stdout by default)
}
```

These are its bones! I think that most of the fields here are self-explanatory,
so I will concentrate on the ones that may not be so obvious...

`Interpreter.running` is a variable responsible for switching interpreter on and
off. `Interpreter.err` obviously contains a `*primitives.Error` and is created
to track errors that arise while the code is being executed. See
[error.go](../DataTypes/Primitives/error.go) to learn more about the `Error`
data type and the way it works.

The two `Environment`s that follow will be ignored for now as they are discussed
in the section below.

The `nil` value in Scoops is just a pointer to the `Nil struct` which holds a
single byte. Separate `struct` is required in order to print and clone `nil`
values correctly. `Nil struct` fully implements the `shared.Object interface`
just like all the other `struct`s in the *Primitives* folder.

And it is only reasonable that we have two pointers to `Writer`s here 
representing the standard output and the current writer (in case user ever
wants to write something to a file for example). There is an interesting benefit
of using `Writer`s instead of `fmt.Printf` -- you can read about it in the
[Shared/types.go](../Shared/types.go).


### Environment type

```go
type Environment struct {
        name    string
        ip      uint64  // instruction pointer
        code    []*shared.Instruction
        data    *stack.Stack
        vars    []shared.Object
        prev    *Environment
}
```

`Environment`s allow *Scoops Programming Language* to execute 
[Scoops](../DataTypes/Scoop/README.md) in separate scopes with their own `data`
stacks, and variable lists.

Most of the fields here are **really self-explanatory** as you can also see
their underlying types.

Just in case you struggle... Instruction pointer is connected to the `code`
slice -- it is literally the **index** of the current `Instruction` that has to
be executed by the *Virtual Machine*.

Every `Environment` also has a pointer to its parent (`prev *Environment`).
Using that pointer, `Environment`s can *receive* arguments from their parents 
and *return* values back.