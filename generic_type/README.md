# Generics in golang 

Generics means parameterized types. Put simply, generics allow programmers to write code where the type can be specified later because the type isn't immediately relevant.

## Examples 

1. [Program without using Generic](wihtout_generic.go)
2. [Using interface type ](interface.go)
3. [Using any type](any_type.go)
4. [Custom Constraints](custom_constraints.go)
5. [Golang Constraints Package](constriant_pkg.go)


## When to use generics

* Functions that operate on arrays, slices, maps, and channels.
* General purpose data structures like stack or linked list.
* To reduce code duplication.
* It is advised to start simple and only write generic code once we have written very similar code at least 2 or 3 times.