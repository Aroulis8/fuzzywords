fuzzywords is a go module that lets you "play around" with words. It includes a variety (I hope!) of functions that let you do exactly that. 

# Functions

## Fuzz
Returns a random sentence made with the words from str

Examples:
_Results may be different due to the randomness of the function_

```go

fuzzywords.Fuzz("Hello World!")

//Returns:
//Hello Hello

```

This function is better with big text

## Reverse
Returns the string you gave it but reversed


```go
fuzzywords.Reverse("Hello World!")

//Returns:
//!dlroW olleH
```