# 99 Problems in Go

This project is an attempt to implement solutions to the classic "99 Problems"
in Go. This set of problems was initially developed for Prolog, and has since
become a popular challenge, particularly for functional programming languages.
Of course, most of them may be implemented in any language, and most attempts
tackle 88 of the problems.

There are many statements of the problems, unless otherwise noted, I follow the
examples in the Haskell wiki since they are particularly clear.

A running list of topics which I have explored in the process of writing these
solutions:

1. Go generic function signatures.
2. Functional Go style which returns copies of slices instead of modifying
   in-place.
3. Generic structs to approximate tuples and nested slices, neither of which
   are built-in to Go.

# Links

- Haskell: https://wiki.haskell.org/index.php?title=H-99:_Ninety-Nine_Haskell_Problems
- OCaml: https://ocaml.org/exercises
- Prolog: https://web.archive.org/web/20170324220754/https://sites.google.com/site/prologsite/prolog-problems
