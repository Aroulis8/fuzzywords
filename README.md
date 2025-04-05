fuzzywords is a go module that lets you "play around" with words. It includes a variety (I hope!) of functions that let you do exactly that. 

# Functions

## Fuzz(str string)
Returns a random sentence made with the words from str

_Results may/will be different due to the randomness of the function_

```go

fuzzywords.Fuzz("Hello World!")

//Returns:
//Hello Hello

```

This function is better with big text

## Reverse(str string)
Returns the string you gave it but reversed


```go
fuzzywords.Reverse("Hello World!")

//Returns:
//!dlroW olleH
```

## RandomCharacters(n int)
Returns a string of "n" random characters

_Results may/will be different due to the randomness of the function_
```go
    fuzzywords.RandomCharacters(10)
    //Returns:
    //pxzczqlyzy
```

If you want it to only use specific characters, you can change randomCharactersSlice
```go
    randomCharactersSlice = []string {"a","b","c"}
    fuzzywords.RandomCharacters(10)
    //Returns:
    //cbccbacbca
```
