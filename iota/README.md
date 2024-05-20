
## iota 
The “iota” identifier is used to represent integer based constants in Go, and is a convenient way to declare a sequence of constants, while keeping the code readable

## Examples 
- [Normal Constant](1_normal_constant.go)
- [Multiple IOTA](2_multiple_iota.go)
- [Single IOTA](3_single_iota.go)
- [Start Non-Zero](4_Non_Zero_Start.go)
- [Skip in between](5_Skiping_Values.go)
- [Custom Math Operation](6_Custom_Math_Op.go)
- [Left Shift](7_left_shift.go)
- [Reset IOTA](8_reset_itoa.go)

## Why Is Iota Useful

Using iota is a more readable and idiomatic way to declare arbitrary constants or enums.
Using iota, we can add more constant values without thinking about the next numerical value to assign.
It also helps us avoid certain types of errors in our code.

To illustrate, imagine we had a large list of error codes to declare as constants. In these cases, its easy to make the human error of assigning the same value twice:
```
const (
	ERR_ABC = 1
	ERR_DEF = 2
	ERR_HIJ = 3
	ERR_KLM = 3 // oops
	ERR_NOP = 4
	ERR_QRS = 5
)
```
If we used iota to assign these values instead, we wouldn’t have this problem, and our code would be less error prone.

## When to Avoid Iota 

- [ ] Iota is best used when there are arbitrary, or incremental values to be used as constants.
- [ ] There isn’t any sequential relationship between constants. For example, time units like second, minute, hour and day can be represented as non-sequential constants:
```
const (
	SECOND = 1
	MINUTE = 60 * SECOND
	HOUR   = 60 * MINUTE
	DAY    = 24 * HOUR
)
```
In this case, there is no way we can use iota because the relationship between successive constants does not increment by the same quantity.

- [ ] Another case where we should not use iota is when the numerical values are not arbitrary. For example, representing the total cards in a deck, or total number of jokers:

```
const (
  TOTAL_CARDS = 52
  TOTAL_JOKERS = 2
)
```

the values of the constants are meaningful, and not sequential, so it’s better to not use iota