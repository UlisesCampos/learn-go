# Is Go an object-oriented language?
Yes and no. Althoigh go has types and methods and allows an object-oriented style of programming, there
is no type hierarchy. The concept of "interface" in go provides a different approach that we believe is easy to use and in some ways more general. There are also ways to embed types in other types to provide something analogous
but not identical—to subclassing. Moreover, methods in Go are more general than in C++ or Java: they can be defined for any sort of data, even built-in types such as plain, “unboxed” integers. They are not restricted to structs (classes).

Also, the lack of a type hierarchy makes “objects” in Go feel much more lightweight than in languages such as C++ or Java.

[docs](https://golang.org/doc/faq#Is_Go_an_object-oriented_language)
