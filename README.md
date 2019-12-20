# Scoops Programming Language 🍨

Consistent and clear syntax. Open source code that is easy to modify. Extensive documentation.

## Motivation

### Abstract

I want to create a high level programming language that allows programmers to *get the job done* easily, that fixes all the things that make you and me go *"uuugh, why is this even a thing!"* when we stumble across some weirdly implemented feature or simply a bad design decision.

### Critique

Most of the popular programming languages have a very decent documentation when it comes to standard library but *poor to no documentation concerning the design of their compiler and/or interpreter*. This significantly narrows the pool of specialists who would be able to contribute to that language not only because they have to be highly skilled but because they are practically forced to sit down and try to recollect all the directory hierarchy of the project (which file is responsible for what, what files are included where, how this thing is built and compiled, and so on) and not every programmer would be thrilled to do that (in fact, I believe, **most of us** wouldn't be).

Moreover, oftentimes the code (and especially the core components) are very unclear which makes it virtually **impossible** to understand, edit, improve or even study what it does and how it does it. For example, try to go through the *"ceval.c"* file of the Python interpreter without shooting yourself or committing violent acts towards a random software developer who was peacefully sipping his coffee on the bench.

Programming language design in itself is an *undecidable problem* as there is no correct answer as to how to manufacture a perfect syntax or a perfect instruction set for the virtual machine. This creates an evolutionary struggle between all programming languages -- they must demonstrate their value on a daily basis to prolong their existence -- and so, just like any other evolution-created things, even the most popular and loved programming languages have their flaws that I want to address.

Think about...

+ the flawed way Python3 handles UTF-8;
+ how you are not really sure whether `a` is just a copy of `b` or is currently pointing to `b`'s original PyObject when you do `a = b` in Python;
+ the vastness of operator overloads (in C++ primarily but also in other languages) that confuse and make code harder to read instead of hiding complexity away as intended;
+ the overcomplicated data types and data structure names in C++ where you see error messages like `std::vector<double>::numba_iterator aka numba_mamba_iterator aka simple_iterator cannot be converted to std::vector<int>::inta_iterator aka inta_pinta_iterator sorry :(`;
+ the vulnerabilities of C's uncontrolled environment and overcomplicated nature of C++'s syntax;
+ the mess of `include`s created by C/C++ header files;
+ the unreasonably globalistic object oriented approach adopted by Java;

### Conclusions

I want to make sure you understand, that one of the main concerns of this project is *simplicity*. I strongly believe that genius things and ideas are **simple** (at least for those who dare to be genius). Thus, complexity blown out of proportion for the sake of complexity is not the way to build a product that people would want to use. I hope you do agree with me here. This, however, does not mean that there won't be any *clever* new things in here -- it only means that those things must *prove themselves to be useful* before joining the family.

Another important aspect that I would like to focus on is *openness*. In all its glory. From open source nature of this project to willingness to write code that is made to be **understood**. Openness means **invitation** to study, share, and contribute. It also means that language should allow for **some** lower level details to be exposed to the programmer while keeping other things hidden if they are important for the performance of the whole system.

Previous two principles rely heavily on the high level of *code readability*. It is crucial to allow my fellow coders to be able to understand how this thing actually works, so it seems reasonable to write a compiler and interpreter that has *extensive and clear documentation* and is written with readability in mind, avoiding overengineered solutions that only add complexity to the overall system.

*General consistency* is another feature to look for in a well designed programming language. It is apparent that there must be more rules than exceptions in a system that dares to call itself reliable, stable, and secure.

---

With ❤️ for 🍨 by [Viktor A. Rozenko Voitenko](https://github.com/sharpvik).
