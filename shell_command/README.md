
## Executing Shell Commands

## Agenda 
* How to execute shell commands (like ls, mkdir or grep) in Golang.
* How to pass I/O to a running command through stdin and stdout, as well as manage long running commands.

## The Exec Package

* We can use the official os/exec package to run external commands.
* When we execute shell commands, we are running code outside of our Go application. In order to do this, we need to run these commands in a child process.
* Each command is run as a child process within the running Go application, and exposes Stdin and Stdout attributes that we can use to read and write data from the process.

## Examples 
1. [Running Basic Shell Commands](1_basic.go )
2. [Executing Long Running Commands](2_long_running_cmd.go)
3. [Custom Output Writer](3_custom_writer.go)
4. [Passing Input To Commands With STDIN](4_parsing.go)
5. [Killing a Child Process](5_kill.go)

## Conclusion
+ So far, we learned multiple ways to execute and interact with unix shell commands. Here are some things to keep in mind when using the os/exec package:

+ Use cmd.Output when you want to execute simple commands that don’t usually give too much output
For functions with continuous or long-running output, you should use cmd.Run and interact with the command using cmd.Stdout and cmd.Stdin

+ In production applications, its useful to keep a timeout and kill a process if it isn’t responding for a given time. We can send termination commands using context cancellation.

## References:





