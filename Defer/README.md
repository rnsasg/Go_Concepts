# When Should You Defer? 

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
1. https://www.sohamkamani.com/golang/defer/