fuzzywords is a go module that lets you "play around" with words. It includes a variety (I hope!) of functions that let you do exactly that.

# Functions

## Fuzz(str string) string

Returns a random sentence made with the words from str

**_Results may/will be different due to the randomness of the function_**

```go
fuzzywords.Fuzz("Hello World!")

//Returns:
//Hello Hello
```

This function is better with big text

## Reverse(str string) string

Returns the string you gave it but reversed
(Thanks to [u/Alternative-Ad-5902](https://www.reddit.com/user/Alternative-Ad-5902/) for the code!)

```go
fuzzywords.Reverse("Hello World!")

//Returns:
//!dlroW olleH
```

## RandomCharacters(n int) string

Returns a string of "n" random characters

**_Results may/will be different due to the randomness of the function_**

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

## AddPrefix(str string, prefix string) string

Adds a prefix to every word in "str"

```go
fuzzywords.AddPrefix("Hello World", "a")
//This returns:
//"aHello aWorld"
```

## AddSuffix(str string, suffix string) string

Adds a suffix to every word in "str"

```go
fuzzywords.AddSuffix("Hello World", "a")
//This returns:
//"Helloa Worlda"
```

## GetStringData(str string) DataString

Gets the amount of **words**, **spaces** and **characters**(without spaces) and returns all of it in a StringData.
(Thanks to [u/Alternative-Ad-5902](https://www.reddit.com/user/Alternative-Ad-5902/) for suggesting me to change the function so it returns a custom struct!)
```go
    fuzzywords.GetStringData("Hello World")
    //Returns:
    //{2 10 1}
```

## CreateFuzzySentence(name string) string

Creates a 'fuzzy'(weird) sentence using random words from **FuzzySentenceVerbs**, **FuzzySentenceObjects** and **FuzzySentencePlaces**
It firsts adds the _name_, then the **verb**, next the **object**, and finally the **place**.

```go
fuzzywords.CreateFuzzySentence("Viss")
//Returns:
//Viss played with a parrot at home
```

