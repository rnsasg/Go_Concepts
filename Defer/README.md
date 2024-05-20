# When Should You Defer? 

Defer is more like a finally block in languages like Java or Python: something that has to be executed after attempting to run some code – whether or not it succeeds.

This is useful in cases where you have to run some clean-up operation after some code executes. For example:

+ Closing a file after opening it.
+ Performing clean-up actions after running a test.

Additionally, you will have to defer any function that needs to handle panics that happen within the function scope.

The main pitfalls you should take note of are argument evaluation and the order of execution for deferred functions, since these can sometimes be confusing for newcomers.

## References
1. https://www.sohamkamani.com/golang/defer/