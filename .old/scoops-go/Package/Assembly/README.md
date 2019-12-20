# Package assembly

*Scoops Assembly* is a language based on simple and clear instructions that are
meant to be executed by the *Scoops Interpreter*. In this file, you can find all
the specs that you should know if you wish to

1. Contribute to this project
2. Create your own *Scoops* compliant compiler or interpreter
3. Branch this repo and modify some core components

## Instruction format

Assembly instruction consists of two parts:

1. Opcode
2. Operand

Regular expression that matches a valid *Scoops Assembly* instruction:

```
^[A-Z_]+( b[01]+| x[\dA-F]+| \d+| '[\x00-\xFF]'| [A-Z_]+)?\s*$
```

## Opcodes

*Scoops Interpreter (SI)* can support up to 256 different opcodes as it is a
[bytecode] based [virtual machine]. Each opcode is a simple command that *SI*
can understand and execute. An example of a trivial opcode would be the `END 0`
opcode that tells *SI* that program has come to an end and it is time to stop
execution.

Some opcodes may require an operand -- additional parameter needed to make *SI*
more flexible. For example, you may want to `PUSH_CONST 42`. This opcode tells
*SI* to push 42 onto the *data stack* ([stack]) of the currently active
execution environment as a `*Byte`, it may then be used to create an instance of
`*Integer` for example.

You can find the list of all the opcodes used by the *Scoops Interpreter*
[in this file]. Regular expression that matches a valid *SI* opcode is as
follows: `[A-Z_]+`

[bytecode]: https://en.wikipedia.org/wiki/Bytecode
[virtual machine]: https://en.wikipedia.org/wiki/Virtual_machine
[stack]: https://en.wikipedia.org/wiki/Stack_(abstract_data_type)
[in this file]: ../Shared/opcodes.go

## Operands

There are 4 types of operands that *Scoops Assembler (SA)* accepts as
syntactically valid:

| Operand Type | Definitions               | Regular Expression | Example      |
|:-------------|:--------------------------|:-------------------|:-------------|
| Boolean      | 8 bits long booleans      | `b[01]+`           | `b101010`    |
| Hexadecimal  | Hexadecimals up to xFF    | `x[\dA-F]+`        | `x2A`        |
| Integer      | Integers from 0 up to 255 | `\d+`              | `42`         |
| Character    | ASCII characters          | `'[\x00-\xFF]'`    | `'*'`        |
| Opcode       | Scoops Assembly opcodes   | `[A-Z_]+`          | `PUSH_CONST` |

The underlying rule is -- every operand must be *1 byte long*. This way, every
instruction is exactly *16 bits (2 bytes)* which makes it extremely easy to
split raw stream of bytes into slice of *Scoops Interpreter Instructions*.