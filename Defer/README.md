# Defer keyword in Go 

## Introduction 

Go's defer statement schedules a function call (the deferred function) to be run immediately before the function executing the defer returns. It's an unusual but effective way to deal with situations such as resources that must be released regardless of which path a function takes to return. The canonical examples are unlocking a mutex or closing a file. 

## Example 1 : The canonical examples are unlocking a mutex or closing a file. 

Deferring a call to a function such as Close has two advantages. First, it guarantees that you will never forget to close the file, a mistake that's easy to make if you later edit the function to add a new return path. Second, it means that the close sits near the open, which is much clearer than placing it at the end of the function.

```go
func Contents(filename string) (string, error) {
    f, err := os.Open(filename)
    if err != nil {
        return "", err
    }
    defer f.Close()  // f.Close will run when we're finished.

    var result []byte
    buf := make([]byte, 100)
    for {
        n, err := f.Read(buf[0:])
        result = append(result, buf[0:n]...) // append is discussed later.
        if err != nil {
            if err == io.EOF {
                break
            }
            return "", err  // f will be closed if we return here.
        }
    }
    return string(result), nil // f will be closed if we return here.
}
```
## Argumnet Evaluations

The arguments to the deferred function (which include the receiver if the function is a method) are evaluated when the defer executes, not when the call executes. Besides avoiding worries about variables changing values as the function executes, this means that a single deferred call site can defer multiple function executions. Here's a silly example.

```go
for i := 0; i < 5; i++ {
    defer fmt.Printf("%d ", i)
}
```
Deferred functions are executed in LIFO order, so this code will cause 4 3 2 1 0 to be printed when the function returns. 

### Another Examples 

```go
func trace(s string) string {
    fmt.Println("entering:", s)
    return s
}

func un(s string) {
    fmt.Println("leaving:", s)
}

func a() {
    defer un(trace("a"))
    fmt.Println("in a")
}

func b() {
    defer un(trace("b"))
    fmt.Println("in b")
    a()
}

func main() {
    b()
}
```

> [NOTE!]
> prints
```text
entering: b
in b
entering: a
in a
leaving: a
leaving: b
```

## When Should You Defer? 

Defer is more like a finally block in languages like Java or Python: something that has to be executed after attempting to run some code – whether or not it succeeds.

This is useful in cases where you have to run some clean-up operation after some code executes. For example:

+ Closing a file after opening it.
+ Performing clean-up actions after running a test.

Additionally, you will have to defer any function that needs to handle panics that happen within the function scope.

The main pitfalls you should take note of are argument evaluation and the order of execution for deferred functions, since these can sometimes be confusing for newcomers.

## Agenda 
1. [Basic Usage](1_defer_simple.go)
2. [Function Scope](2_defer_function_scope.go)
3. [Execution Stack](3_defer_execution_stack.go)
4. [Defer with Panic](4_defer_with_panic.go)
5. [Defer Recover over Panic](5_defer_panic_recover.go)
6. [Argument_Validation](6_defer_argument_evalution.go)
7. [Defer_Output](7_defer_outout.go)

## References
1. [Defer Blog](https://www.sohamkamani.com/golang/defer/)
2. [Effective Go](https://go.dev/doc/effective_go#defer)