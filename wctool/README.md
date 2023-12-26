## Setup project

To create a Go project, we will first verify Go is installed

```sh
go version
```

Then create the Go project.

```sh
mkdir wctool
cd wctool 
go init vigneshm243/wctool
```

## Implementing the functionalities

I will be honest I didn't do this challenge linearly as I didn't want to refactor the code for any surprises that might come later. So, I went through the entire problem in one go and started with the challenge. On second thought, I should have gone through it linearly to learn if I am actually writing code that can handle such suprises.

So the main objective of the progam is to find the number of words, characters, bytes, lines in a file or from std input. 

Flags supported for the command we are creating

| Flag | Operation |
|---|---|
| c | Count bytes |
| l | Count lines |
| w | Count words |
| m | Count characters |

I have written only 2 files for the project they are main.go and wordcount.go

main.go - Take the input, create a Reader and call the CalculateWordCount
wordcount.go - To do the actual calculation of the stats and a simple function to format the output

To build we run in the wctool folder

```sh
go build .
```

This creates a wctool.exe or wctool depending on the OS. Then we can execute it as we would execute any CLI tool

```sh
wctool test.txt
```

## Some Learnings and Discoveries

Having not used Go much, these might be pretty low level, but I found them interesting enough to jot em down.

- ### Flag package
    Go has a built in flag package which can be used for flag parsing. A simple usage of it as follows,

    ```go
    var countBytes bool
    flag.BoolVar(&countBytes, "c", false, "Count bytes")
    ```
    Let's break down this simple line.

    - flag  - the package (imported in the file)
    - BoolVar - type of the flag
    - &countBytes - variable to bind to
    - "c" - the flag that will be used in the command
    - false - default value
    - "Count Bytes" - the flag message

- ### Read args from CLI

    We can read from the CLI simply in Go by using the os package Args variable.

    ```go
    args := os.Args
    ```

    A peculiar thing to note here is unlike other programming languages Go seems to include the program in the list of args. This is stored in the 0th index. So actual args start from 1 which is a slice of strings

    ```go
    programName := os.Args[0] //the program name
    argsWithoutProgram := os.Args[1:] //slice of args without the program name
    ```

    In a program where we are using flag package it's better to use the flag.CommandLine.Args. This should be called after flag.Parse() else will result in the flags being considered as args.

    ```go
    flag.Parse() // parses the flags and assigns the values to referenced vars
    fileNames := flag.CommandLine.Args() // read the list of filenames after the flags have been passed
    ```

- ### Read File Input and read Console Input
    - #### File Input
    
        In Go, you can read from a file in various ways:

        - Using os.ReadFile() to read an entire file at once.
        - Using bufio.NewScanner() to read a file line by line.
        - Using bufio.NewReader() to read a file in buffered way.

        In our program we have used NewReader and read the file Rune by Rune.

        To open a file we use the os package.

        ```go
        file, err := os.Open("file.txt")
        if err != nil {
            log.Fatal(err)
        }
        defer file.Close()
        ``` 
        Also, note here we are using defer file.Close() which schedules the close of the file. This is to avoid resource leaks. This will be called when the main function finishes execution, whether normally or due to an error. (Similar to a finally block in Java maybe)

    - #### Std Input

        In Go, you can read from the console using several methods:

        - Use bufio.NewReader().ReadString() to read a full line of text, including the newline character.
        - Use bufio.NewScanner().Scan() to read a full line of text without the newline character.
        - Use bufio.NewReader().ReadRune() to read a single character.
        - Use fmt.Scanln() to read each word of a line into a different variable.

        In our program we have used NewReader to create a buffered reader which is passed to the same function that handles both file and std in.

- ### Rune
    A rune is an alias for int32. It holds a single unicode character.

- ### Difference between Reader and Scanner
    
    In Go, Reader and Scanner are part of the bufio package. Reader is a low-level interface that reads raw bytes into a buffer, providing flexibility for various input sources. Scanner is a higher-level interface that simplifies reading values of different types from the input, automatically handling conversion of input into basic data types. Performance-wise, Reader can be more efficient for optimized read operations, while Scanner can be slower and allocate more memory when reading large amounts of data.

- ### Naming conventions in Go
    In Go, naming conventions affect the visibility and usability of identifiers. 

    - Identifiers starting with an uppercase letter are exported (visible outside the package), while those starting with a lowercase letter are not exported (only visible within the package). 
    - Multiword names are written in MixedCaps or mixedCaps (camelCase), not underscores. 
    - Package names are short, clear, lowercase, and usually simple nouns. 
    - Function and variable names should be descriptive and clearly indicate their purpose. 

- ### Optional Design pattern

    The optional design pattern, also known as the functional options pattern, can be used in Go to provide a flexible way to pass optional configuration parameters. 

    Here we have used by the use of commandLineOptions and implemented it with the help of the flag package. Let's understand a little about it and the advantages it provides.

    Here we have defined a Options struct that has the list of configuration flags we will be using in this command. 

    ```go
    type Options struct {
        countBytes bool
        countLines bool
        countWords bool
        countChars bool
    }
    ```

    Later if we want to add more options we can modify this struct to do so, rather than rewrite more code if we had used seperate variables for each of the option. 

    ```go
    var commandLineOptions Options
    flag.BoolVar(&commandLineOptions.countBytes, "c", false, "Count bytes")
    ```

    Here we are defining a single Options variable called commandLineOptions that is being passed around functions. Also, we are defining the default values for these options if the user doesnot provide them.

### Things I want to add to this 

- Parallelize the reads by chunking the file and process simultaneously using go routines
- Check the performance and memory usage of the progam and compare it with the actual wc implementation
<!-- Time program to check if program is fast enough -->
<!-- Buffered reader / Parallel Streams -->