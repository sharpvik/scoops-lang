# Scoops Programming Language 🍨

Consistent and clear syntax. Open source code that is easy to modify.
Extensive documentation.



## Table of Contents

1. [Overview](#overview)
2. [Critique & Motivation](#critique-and-motivation)
3. [Aims](#aims)
4. [User Guide](#user-guide)
    - [Scoops Execution Suite](#user-guide-scoops)
    - [Scoopcy Compilation Suite](#user-guide-scoopcy)
5. [Contributors](#contributors)



## <a name="overview"></a> 1. Overview

*Scoops* is a high level programming language that allows programmers to
*get the job done* easily, that fixes all the things that make you and me go
*"uuugh, why is this even a thing!"* when we stumble across some weirdly
implemented feature or simply a bad design decision.



## <a name="critique-and-motivation"></a> 2. Critique & Motivation

Most of the popular programming languages have a very decent documentation when
it comes to standard library but *poor to no documentation concerning the design
of their compiler and/or interpreter*. This significantly narrows the pool of
specialists who would be able to contribute to that language not only because
they have to be highly skilled but because they are practically forced to sit
down and try to recollect all the directory hierarchy of the project (which file
is responsible for what, what files are included where, how this thing is built
and compiled, and so on) and not every programmer would be thrilled to do that
(in fact, I believe, **most of us** wouldn't be).

The *Scoops Programming Language* is supposed to be **easily understandable**.
This is achieved using extensive documentation, not only of standard library and
high level syntax, but also of its compiler's and interpreter's core components.
This allows *Scoops* to be *great for education* purposes, open to 
*third party customisation*, *easy to learn* and use.



## <a name="aims"></a> 3. Aims

One of the main concerns of this project is *simplicity*. We strongly believe
that genius things are **simple** (at least for those who dare to be genius).
Thus, complexity blown out of proportion for the sake of complexity is not the
way to build a product that people would want to use.

This, however, does not mean that there won't be any *clever* new things -- it 
only means that those things must *prove themselves to be useful* before joining
the family.

Another important aspect is *openness*. In all its glory. From open source 
nature of this project to willingness to write code that is made to be 
**understood**. Openness means invitation to *study, share, and contribute*.

*Consistency* is another feature to look for in a well designed programming
language. It is apparent that there must be more rules than exceptions in a
system that dares to call itself reliable, stable, and secure.



## <a name="user-guide"></a> 4. User Guide

This project contains two different tools:

1. *Scoops* -- Execution Suite subdivided into two implementations:
    - [scoops-rs] implemented in Rust
    - [scoops-go] implemented in Go
2. *[Scoopcy]* -- Compilation Suite

[scoops-rs]: scoops-rs
[scoops-go]: https://github.com/sharpvik/scoops-go
[Scoopcy]: scoopcy


### <a name="user-guide-scoops"></a> Scoops Execution Suite

***Scoops*** is a program that takes your script and produces desired output.
If you wish to execute your *Scoops* code, pick one implementation, click on the
link and follow instructions in `README.md` file.


### <a name="user-guide-scoopcy"></a> Scoopcy Compilation Suite

***Scoopcy*** is a program that accepts `*.scp` and `*.scpa` files and is able
to produce `*.scpa` and `*.scpb` files, compiling and assembling your source
code.

| Source   | Can Be Compiled To |
|:--------:|:-------------------|
| `*.scp`  | `*.scpa`, `*.scpb` |
| `*.scpa` | `*.scpb`           |

*Scoopcy* can be used to produce intermediate *Scoops* code forms such as
*Scoops Assembly* and *Scoops Bytecode* which start executing faster as they
don't require syntax and semantics to be checked. *Scoopcy* is also a great
tool that allows you to study and understand how *Scoops* actually works.

Using *Scoopcy* is easy and straightforward -- you can find instructions
[here](scoopcy/README.md).



---

## <a name="contributors"></a> Created with ❤️ for 🍨 by:

- [Viktor A. Rozenko Voitenko](https://github.com/sharpvik)
- [Konrad Krzysztof Sobczak](https://github.com/kondziusob)
