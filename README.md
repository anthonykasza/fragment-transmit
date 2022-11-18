Some Notes on Go
================

packages - namespaces for organizing related functions and globals. the `package` keyword is similar to Zeek's `module` and `export` declarations.

exported functions - begin with capital letters

error handling - go doesn't have `try` and `catch` like Python does. many functions will return a value or an error like `return DateTime, nil`. similarly, assignment looks like `myDateTimeValue, nil := package.func()`. this is similar to Solidity (and somewhat reminds me of Rust's unwrap() functionality). be sure to check and handle conditions of errors.

goroutines - `go func()`. they keyword `go` makes a function execute in a thread. goroutines access shared memory synchronously.

channels -  `chan` keyword. how goroutines send/recieve messages. channels are powerful in Go and are the cornerstone to Go's concurrency model. channels are of some type, similar to how Zeek's opaques are "of" some type. channels can be buffered. channels can be closed by a sender and checked by a receiver ala, `v, ok := <-ch`. closing a channel is optional and really only useful when telling a `range` statement there is no more dta coming. channels can also have a direction attribute associated with them when passed as func arguments, for example `func (messages chan<- string) {}`.
channels are very similar to how TCP works. the sending endpoint SYNs while the receiving endpoint ACKs incoming segements (Go's "segments" are typed). For bidirectional communications TCP's receiver will also SYN back to the sender (as part of the SYN+ACK) and the sender will ACK (the 3rd step of the 3-way handshake).
ONLY SENDERS SHALL CLOSE A CHANNEL. 
CHANNELS SHALL BE UNIDIRECTIONAL UNFORCED BY ATTRIBUTES.

select - a statement that lets a goroutine wait on multiple channels. `select` blocks until one of it's `case`s can run - sort of like a `switch` but channel driven.

types hints - go is strongly typed but the syntax supports implicit typing. a list of variables with all the same type may have only 1 type hint. for example `var i, n, j int = 1, 99, 2`. compare to `var i int, n uint64, j string = 1, 99, "2"`

atomic types - bool, string, int[8,16,32,64], uint[8,16,32,64], uintptr, byte (alias for uint8), rune (alias for int32 and use for unicode), float[32,64], complex[64,128]

pointers - go's gottem. `var p *int` declares a a pointer. `p := &i` says that `p`'s value points to the address which `i` lives at. unlike C, Go doesn't require us to do pointer arithmetic by hand, hurray! DONT USE POINTERS FOR REFERENCE TYPES, like channels. DONT NEST POINTERS, gross. 

reference type - a variable that references another here be dragons: https://github.com/go101/go101/wiki/About-the-terminology-%22reference-type%22-in-Go

value type - a variable that represents a value. here be dragons: https://github.com/go101/go101/wiki/About-the-terminology-%22reference-type%22-in-Go

structs - similar to C structs or records in Zeek. referencing a field of a struct is done with a period. structs can be referenced using pointers. the following are equivalent, `(*p).X` and `p.X`. fields in structs have default initialization values.

arrays - array expressions look weird, `var a [10]int` makes an arrary of 10 integers. `a := [...]int{0,1,2,3}` makes an array of 4 integers. arrays can be sliced, just like in Python or Zeek.

slices - a slice is a reference to a contiguous segment of an array. arrays are value types, slices are reference types. slices have a len() and capacity, cap(). slices are types in Go, unlike in Python and Zeek where slicing is an operation instead of a type. the "zero value" of a slice is `nil`. slices support appending with `s.append()`. this increases the capacity and length of the slice by 1 but the original array/slice the slice references does NOT get appended to.

maps - similar to Python's `dict`s and Zeek's `table`s. maps default "zero value" is `nil`. missing key errors are found with something like, `v, err := m["missing"]`

functions - are types too, `fmt.Println(func_name)` will result in a pointer being printed. functions can be closures

default initialization - types have default "zero" values. 0, false, "" are default values. for example, `var b int` sets `b` equal to 0.

null - nil, "(<nil>)"

typecasting - call the type as function, `f = float64(i)`. types are inferred by default

variable declaration - there's an explicit version, `var foo int = 4` and a shorthand version, `foo := 4`. `:=` is a declaration and assignment. `=` is an assignment only. `var foo int` is a declaration. like Solidity, you can declare and assign multiple variables in a single statement, like `var i, s := 1, "str"`. 

naked return statements - gross, i don't like them. they are implicit return statements where the returned variables are matched, by name, in the declaration and body of the function. gross. 

loops - go uses `for` loops. they behave like regular for-loops and also like while-loops. 1 type of loop is nice, compare with Perl's mess of `do`, `while`, `foreach`, `until`, and more... however, Go has loop modifiers like `range`. `continue` and `break` behave similarly to other C-based languages

range - a modifier to the `for` loop which combines with a slice or map to return 2 values per iteration. the first is the index and the second is a copy of the element.

conditions - go allows for variable declaration in `if` statements' conditions. those variables are scoped to the if's body. `if v := math.Pow(x, n); v < lim {return v}`. `if`s can be paired with `else`s. go supports `switch` statements too which means `case` and `default` are reserved keywords. 

defer - a function invocation modifier that ensures the function is executed sometime after it's scope ends. defer functions in a single body are handled LIFO (stack) ordering.

go - another function invocation modifier. the keyword `go` indicates the function should be called in a thread.

make() - a built in function that returns a new slice, map, or chan. `make` behaves 3 different ways depending on what you're making.

_ - similar to Solidity, this is the common name of a "throw away" value.

methods - go does not support classes but types can "receive" methods. to declare a method on a type, use a `func` declaration with a special "receiver" argument. see methods.go for examples of when to use a receiver value or receive reference (hint: when in doubt us a reference)

interface - a set of method signatures. these are similar to interface types in TypeScript except they only contain methods. an `interface` is a list of methods a value must implement. if a value implements all of those methods (by name/argument/return type signature) then it is considered of that interface type. calling a method on an interface value calls the method on the underlying type. different modules implement different interfaces. for example, `String` implements the `Stringer` interface for anything that can be described as a string

empty interface - `interface{}`. an interface with zero methods. every value implements at least zero methods, therefore every value matches an empty interface. empty interfaces are used when value's types are unknown (this is how you break out of strong typing). casting an empty interface to a type can be done with a type assertion, `var i interface{} = "hello"; s := i.(string)`. switch statements are handy for handling empty interface types.

goto - the `goto` keyword is used to manipulate the procedural flow of a func

iota - something to do with numbering things... i'm not sure

errors - a value that represents an error. errors are built-in interfaces

readers - another common interface "which represents the read end of a stream of data". 

[commonlibs](https://pkg.go.dev/std)
-----------
fmt
time
reflect
sync
math
os
net
errors


thanks to
---------
- https://gobyexample.com/
- https://go.dev/tour
- https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/ch04.html
