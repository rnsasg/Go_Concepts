## What are function closures in Go? [link](../function/closure.go)

Closure --> Anonymous functions.

Closures are functions that capture and use variables from their surrounding lexical scope.
They are often used in Go to create anonymous functions.

Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

```
package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

// Output:
// 0 0
// 1 -2
// 3 -6
// 6 -12
// 10 -20
// 15 -30
// 21 -42
// 28 -56
// 36 -72
// 45 -90
```
## Use Cases of Function Closures in Go

### Encapsulation of State:
Closures can encapsulate state that persists across multiple function calls. This is useful for maintaining state in a function without exposing it globally.

```
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    count := counter()
    fmt.Println(count()) // 1
    fmt.Println(count()) // 2
}
```
### Callbacks and Event Handlers:

Closures are often used as callbacks or event handlers where the function needs to access or modify the state that is defined outside of its scope.

```
func createCallbackFunction(base int) func(int) int {
    return func(i int) int {
        return base + i
    }
}

func main() {
    addFive := createCallbackFunction(5)
    fmt.Println(addFive(3)) // 8
}
```

### Factory Functions:

Closures can be used to generate functions with specific configurations, essentially acting as factory functions.

```
func multiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

func main() {
    double := multiplier(2)
    triple := multiplier(3)
    fmt.Println(double(4)) // 8
    fmt.Println(triple(4)) // 12
}
```

### Iterators and Generators:

Closures can maintain internal state for iterators and generators, providing a way to yield values over time.

```
func integerGenerator() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {
    nextInt := integerGenerator()
    fmt.Println(nextInt()) // 1
    fmt.Println(nextInt()) // 2
}
```
## Benefits Over Normal Functions

### State Maintenance:

Closures can maintain state between function calls, which normal functions without closures cannot do without global variables or passing state explicitly.

### Data Hiding and Encapsulation:

Closures enable data hiding, as the state maintained within a closure is not accessible from outside the function. This provides a way to encapsulate private data.

### Higher-Order Functions:

Closures are useful in higher-order functions where functions are passed as arguments or returned from other functions, allowing for more flexible and reusable code.

## Cleaner Code:

Using closures can lead to cleaner and more readable code by reducing the need for external state variables and encapsulating functionality within functions.
When to Use Closures

## When Encapsulating State:

Use closures when you need to maintain state across multiple calls to a function without using global variables.

## When Implementing Callbacks:

Closures are ideal for callbacks where you need to capture the state at the time the callback is created.

## When Creating Factory Functions:

Use closures to create functions with pre-configured parameters, enhancing code reusability.

## When Designing Iterators:

Closures are suitable for designing iterators that need to maintain internal state between successive iterations.
In summary, closures in Go provide a way to encapsulate state, hide data, and create more flexible and reusable code structures. They are particularly useful in scenarios requiring state persistence, callbacks, factory functions, and iterators.







