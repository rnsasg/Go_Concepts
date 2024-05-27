# Using Maps in Golang

Maps are a built-in data structure in Go that act as a collection of key-value pairs. They are also known as associative arrays, hash tables, or dictionaries in other programming languages.

## Creating a Map

### Declaration 
When we want to create a new map, we need to declare its key and value types. The syntax for declaring a map type is:
```
var m map[keyType]valueType
```
This declares m as a map with keys of type `keyType` and values of type `valueType`

### Creating using make 

Above examples only declares the type, but to create an actual map, we need to use the built-in make function.
This creates an empty map with `string` keys and `int` values.
```
myMap := make(map[string]int)
```

### Adding Key-Value Pairs to a Map
We can add key-value pairs to a map using the following syntax:
```
m[key] = value
```
This adds a key-value pair to the map m. If the key already exists in the map, then the value is updated. If the key does not exist, then a new key-value pair is added to the map

- [ ] example,if we want to add the key-value pair "the_answer": 42 to the map myMap, we can do so as follows:
```
myMap["the_answer"] = 42
```
If you set a key that already exists, then it’s previous value is overwritten.

### Getting a Value from a Map

We can get the value associated with a key from a map using the following syntax:
```
value = m[key]
```
This returns the value associated with the key key in the map m. If the key does not exist in the map, then the zero value of the value type is returned.

Let’s see an example of this. If we want to get the value associated with the key "the_answer" from the map myMap, we can do so as follows:
```
value := myMap["the_answer"]
```

However, if we try to get the value associated with the key "the_question" from the map myMap, we will get the zero value of the value type, which is 0 in this case.
```
value := myMap["the_question"]
// value == 0
```

### Checking if a Key Exists in a Map

Since we get a zero value when a key doesn’t exist in a map, we can’t use this to determine if the key was present in the map or not. (What would we do if the key was intentionally set to the zero value?)

Fortunately, when obtaining a value from a map, we can also get a boolean value that indicates whether the key was present in the map or not:
```
value, ok := m[key]
```
The value of the second variable ok is true if the key was present in the map, and false otherwise.

### Deleting Key-Value Pairs
We can delete a key-value pair from a map using the inbuilt delete function:
```
delete(m, key)
```
For example, if we want to delete the key-value pair "the_answer": 42 from the map myMap, we can do so as follows:
```
delete(myMap, "the_answer")
```
The delete function works even if the key doesn’t exist in the map, or even if the map is nil. In that case, it does nothing, and acts as a no-op.

### Iterating Over a Map
We can iterate over a map using the range keyword. In each iteration, we can access a the key and value variables:
```
for key, value := range m {
    // do something with key and value
}
```

> [!NOTE]
> Note that the order of iteration is not guaranteed (even for the same map in two different iterations). If you want to iterate over the map in a specific order, you will need to sort the keys first, or else maintain a separate slice of keys in the order you want.


## Allowed Types for Keys

There are some restrictions on the types that can be used as keys in a map. The key type must be one of the following:

Any type that is comparable using the == operator. This includes all the basic types, such as int, string, bool, etc. and also pointers, channels and interfaces.
Any struct or array type that contains only types that are comparable using the == operator.
Slices, functions, and maps cannot be used as keys, since they are not comparable using the == operator.








