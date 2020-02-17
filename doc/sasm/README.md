# Scoops Assembly -- Architecture & Standards



## Introduction & Motivation

*Scoops Assembly* is a middle-level language that allows people to learn about
the inner structure of *Scoops* and its architecture.

Normally, you'd write all your *Scoops* code using [Scoops Source Code][src] and
use some implementation of *Scoops* (e.g. [scoops-go] or [scoops-rs]) to execute
that code. That's how most of us go about it... but what if you wanted to
understand what *actually* happens under the hood?

[src]: ../synt
[scoops-go]: https://github.com/sharpvik/scoops-go
[scoops-rs]: ../../scoops-rs

Well, in that case you read the [documentation][doc] and play with [scoopcy].
*Scoopcy* can clearly demonstrate how your source code transforms into bytecode
before it gets picked up by the [Scoops Virtual Machine][scoops-rs] and
executed.

[doc]: ../README.md
[scoopcy]: ../../scoopcy

*Scoops Execution Suite* doesn't generate *Scoops Assembly* during compilation
process -- it jumps straight to bytecode generation to decrease execution time,
however, you can understand *Scoops Bytecode* much better and to a greater
extent by learning about *Scoops Assembly* and its architecture as it translates
(almost) one-to-one into bytecode that the *Scoops Virtual Machine* uses.



## Architecture


### Constant Literals

First things first, to increase *Scoops* execution speed, we need to have
sections with **constant literals** that come before instructions.

```scpa
$flt                # <- this hash means 'comment'
3.14, 9.81, 1.618

$byte               # '$' stands for 'section'
97, 98, 99

$str                # after '$' comes appropriate data type
"hello, world", "Scoops is a scope-based, interpreted programming language"
```

It is apparent from the name, that these constants are derived from literals by
the compiler. For example, if you ever say `var a = 42`, that `42` will be put
into `$int 42`.

*Scoops Virtual Machine* will know exactly how to build appropriate data
structures at runtime using that section's name and the data we gave it.
Compiler, sadly, cannot predict the exact binary pattern of a data structure
built from a constant literal, so it is left for the VM to figure it out,
however, by giving that data in such condensed blocks instead of using
instructions to `push`, `pop`, and `make` these data structures on the tip of
the `data stack` we significantly reduce execution time.

Of course, not all types can be put down as a constant -- collections, such as
`List`, `Dict`, and `Set` will have to be constructed using constants and some
stack manipulations.

```scpa
$int
2       # becomes 0th const

$flt
3.14    # 1st const

$str
"four"  # 2nd ... you get it :)

$exe    # here, instruction section begins
ldcq 2  # load const quick ('quick' means 'from operand' instead of data stack)
ldcq 1
ldcq 0
mkintq 3    # make int quick
mklist  # omitting the operand means it's not important, inserts 0 by default
print   # print object on top of the stack
pnl     # print newline

# starting cleanup and exit...

pops    # pop data stack
mknil   # make value of `nil` type
pushq   # pop value from the top of the data stack (`nil`), push it into the
        # interchange queue (used for passing arguments to and returning values
        # from scoops)
return  # as soon as the global environment returns, program ends

# return from a normal scoop-environment simply discharges all of the contents
# of its interchange queue into its parent environment's interchange queue
```

Data types that are allowed to become constants:

| Data Type | Example | Description                                            |
|----------:|:-------:|:-------------------------------------------------------|
| `byte`    | `b'a'`  | This data type represents a number between 0 and 255   |
| `flt`     | `3.14`  | This is a floating point number                        |
| `int`     | `42`    | A regular integer                                      |
| `nil`     | `nil`   | The `nil` type has a special semantic significance     |
| `rune`    | `'🍨'`  | One `rune` represents a variable length UTF-8 symbol   |
| `str`     | `"Hi!"` | UTF-8 string that consists of `rune`s                  |

The `bln` (boolean) data type is not listed here since it is too simple to make.
It doesn't need to be prerecorded as it only takes one bytecode instruction to
make it.
