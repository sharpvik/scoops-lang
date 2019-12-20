# Scoops Assembly -- Architecture & Standards



## Introduction & Motivation

*Scoops Assembly* is a middle-level language that allows people to learn about
the inner structure of *Scoops* and its architecture.

Normally, you'd write all
your *Scoops* code using [Scoops Source Code][src] and use some
implementation of *Scoops* (e.g. [scoops-rs]) to
execute that code. That's how most of us go about it... but what if you want to
understand what *actually* happens under the hood?

[src]: ../src/README.md
[scoops-rs]: ../../scoops-rs/README.md

Well, in that case you read the [documentation][doc] and play with [scoopcy].
*Scoopcy* can clearly demonstrate how your source code transforms into bytecode
before it gets picked up by the [Scoops Virtual Machine][scoops-rs] and
executed.

[doc]: ../README.md
[scoopcy]: ../../scoopcy/README.md

*Scoops Execution Suite* doesn't generate *Scoops Assembly* during compilation
process -- it jumps straight to bytecode generation to decrease execution time,
however, you can understand *Scoops Bytecode* much better and to a greater
extent by learning about *Scoops Assembly* and its architecture as it translates
(almost) one-to-one into bytecode that the *Scoops Virtual Machine* uses.
