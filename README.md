# Golang Notes

This repository contains my personal notes, code snippets, and examples as I embark on my journey to learn Go (Golang). The aim is to document key concepts, language features, and practical implementations while building a strong foundation in Go programming.

# Index
- [Variables](#variables)
    - [Formating Variable](#variable-formatting)
- [Arrays and Slices](#arrays-and-slices)
- [Conditions](#conditions)
- [Loops](#loops)
    - [For](#for)
    - [Map](#map)
    - [Foreach](#foreach)
- [Functions](#functions)
- [Pointers](#pointers)
- [Structs](#structs)
- [Interface](#interface)
- [Defer](#defer)
- [Scopes](#scopes)
- [Packages and Modules](#packages-and-modules)
- [Function Literals](#function-literals)
- [Contributing](#contributing)
- [Resources](#resources)

# [Variables](variables/main.go)
This section contains examples and explanations of variable declarations and usage in Go.

The key issue while determinig a variables in Go by using this method;

```go
variablesName := value
```
This method is used just in function

| Feature                  | `var`                                      | `:=`                                        |
|--------------------------|--------------------------------------------|---------------------------------------------|
| **Scope**                | Can be used both inside and outside functions. | Can only be used inside functions.          |
| **Type Declaration**     | May be required (optional).                | Type is automatically inferred.             |
| **Use Cases**            | Used for global variables, constant declarations. | Used for quick and concise declarations.    |
| **Reassignment**         | New values can be assigned.                | Cannot be used for reassigning existing variables. |

- # [Variable Formatting](variables/main.go)

| Feature         | `%s`             | `%d`             | `%f`             | `%t`              |
|-----------------|------------------|------------------|------------------|-------------------|
| **Data Type**   | `String`          | `Integer`        | `Floating-point` | `Boolean`         |
| **Description** | Represents the product name. | Represents the product quantity. | Represents the product discount. | Represents whether the product is in stock. |

# [Arrays and Slices](arrays_slices/main.go)
Arrays are two types in Go: fixed and non-fixed. First of all you have to determine type of the arrays. 

```go
var array = [2]string {"value1", "value2"} ---> its fixed
var array = []int {1,2,3} ---> its non-fixed
```
# [Conditions](conditions/main.go)

The `if-else` statement allows you to define an alternative block of code to execute if the condition is false. Additionally, the `else if` statement enables you to check multiple conditions in sequence, executing different blocks of code based on the first true condition.

```go
if condition1 {
    // code for condition1
} else if condition2 {
    // code for condition2
} else {
    // code if none of the conditions are true
}
```
| Operator  | Description                         | Example                         | Result                                            |
|-----------|-------------------------------------|---------------------------------|---------------------------------------------------|
| **`and`**  | Logical AND                        | `a > 10 && b < 20`              | Returns `true` if both conditions are true.       |
| **`or`**  | Logical OR                          | `a > 10 or < 20`                | Returns `true` if at least one condition is true. |

# [Loops](loops/main.go)
 - # For
This is similar to loops in languages like C, Java, or JavaScript. It has three parts:

 - Initialization: Executed once before the loop starts.
 - Condition: Checked before each iteration. If true, the loop continues; if false, the loop terminates.
 - Post statement: Executed after each iteration.

 ```go
for index:=1; index<=10; index++ {
    fmt.Println(index)
}
```
- # [Map](maps/main.go)

 - Map Declaration: The program defines a map names, where city names (strings) are keys and integers are values. The initial map contains:
"Paris": 0
"NYC": 1
"London": 2
 - Element Deletion: The `delete()` function removes an element by its key. In this case, the "NYC" entry is deleted:
delete(names, "NYC")
 - Printing the Result: After deletion, `fmt.Println()` is used to print the updated map:
`fmt.Println(names)`
 - Output:
`map[Paris:0 London:2]`
The "NYC" entry is removed, leaving only "Paris" and "London".
`delete()` Function: The `delete()` function removes the specified key-value pair from the map. If the key does not exist, it does nothing without producing an error.
 - Summary:
This program shows how to use a map in Go and how to delete an element from it using the `delete()` function. The map is printed after the deletion, showing the updated content.

- # [Foreach](foreach/main.go)
This code demonstrates the use of range in Go to iterate over various data structures. The program performs the following operations:

 - Iterating Over a Slice
The coding_languages slice is iterated over using range, and each programming language is printed.
 - Iterating Over a Map (Values Only)
The dict map is iterated to print only its values.
 - Iterating Over a String
The tech string is iterated character by character using range.
 - Iterating Over a Map (Keys and Values)
The dict map is iterated to print both keys and their corresponding values.

# [Functions](functions/main.go)

 - Function Definition and Calling
Functions are defined using the `func` keyword and perform a specific operation when called. Functions can take parameters and return a value.
 - Parameters and Return Values
Functions can take one or more parameters. A function can return one or more values.
 - Slices and Variable Number of Parameters
Functions can accept slices (arrays) to operate on multiple elements. Functions can also be defined to accept a variable number of parameters using ....

# [Pointers](pointers/main.go)
A pointer is a variable that stores the memory address of another variable.

Key Concepts:

- Declaring a Pointer:
A pointer is declared with the * symbol.
- Example: var `ram_adress` `*int` means `ram_adress` is a pointer to an integer.
- Getting the Address:
The & operator is used to get the memory address of a variable.
Example: `ram_adress` = `&number` stores the address of number in ram_adress.
- Dereferencing a Pointer:
The `*` operator is used to access or modify the value at the memory address.
Example: `*ram_adress` = 20 changes the value of number to 20.

# [Structs](structs/main.go)
In Golang, the struct type is used to hold multiple fields of different types together and is an ideal tool for defining the properties of an object. A struct consists of fields, and each field has a name and a type.

Methods can be defined for these data structures created with struct, allowing operations to be performed on the specific struct type. Methods encapsulate functionalities such as reading or modifying the properties of a struct.

- Value and Pointer Reference:
Methods can operate on a copy of the struct or its reference. If you want to modify the fields of a struct in a method, you need to use a reference (pointer).
- Object-Specific Features:
In Golang, struct provides the basic features of classes in object-oriented programming, but since Go focuses on simplicity, it does not have inheritance. Instead, composition is used to include the features of one struct in another.
- Access Levels:
In Go, fields and methods are accessible from outside their package if their names start with an uppercase letter (public). If they start with a lowercase letter, they are only accessible within the package they are defined in (private).

# [Interface](interface/main.go)

# [Defer](defer_usage/main.go)
In Go, `defer` postpones the execution of a function until the surrounding function finishes. `defer` is ideal for cleanup tasks, like closing files or releasing resources. Deferred calls execute in Last In, First Out (LIFO) order.

# Scopes
In Go, variable scope determines where a variable can be accessed:
- Package Scope:
Variables declared outside functions are accessible throughout the entire package.
- Function Scope: 
Variables declared inside a function are accessible only within that function.
- Block Scope:
Variables declared inside a block (e.g., within if, for, or {}) are only accessible within that block.
- Shadowing:
Inner variables with the same name as outer variables temporarily hide the outer ones within their scope.

# [Packages and Modules](modules/main.go)
1) **Module**\
A module is the basic unit of a Go project, where all the files and packages are linked together. The `go.mod` file located in the project directory is used for managing the module's name and dependencies. A module typically has a general name that spans the entire project, and packages are imported using this name.
`go.mod`file created by: 

```bash
go init moduleName
```

2. **Packages**\
Packages help in organizing the code and structuring different functionalities into separate files. Each `.go` file must belong to a package, which is specified at the beginning of the file. Packages are organized within the module as subdirectories. Each package contains functionalities specific to its purpose. To use a package in another file, it is imported using the `import` statement.

3. **Using Packages**\
To use the functions of a package in another file, the module name along with the package name is specified in the `import` statement. For a function in a package to be accessible externally, its name must start with an uppercase letter. In Go, identifiers that start with an uppercase letter are considered "exported" and can be called from other packages. Once exported, functions and other structures can be accessed using the package name.

4. **Advantages of Package and Module Organization**
- Code Organization: Code is more organized and readable because it is divided according to functionality.
- Reusability: Packages can be easily reused within the same module or in other projects.
- Dependency Management: The `go.mod` file simplifies managing and ensuring version compatibility of dependencies.
- Testability: Independent test files can be created for each package, making the project easier to test.

# [Function Literals](function_literals/main.go)
Function literals, also known as anonymous functions or lambda functions, are functions that are defined without a name. They are often used when you need to pass a function as a parameter, return a function, or define a simple function for immediate use.

1. Definition: A function literal is defined with the `func` keyword, followed by parameters and the function body, but without a name.
2. Usage: Function literals can be assigned to variables, passed as arguments to other functions, or immediately invoked.
3. Syntax: A function literal can accept parameters and return values, just like a regular function.
It is often used when a simple, one-time-use function is needed.
Function literals provide a compact way to define and use functions without the need for a formal function declaration, making them useful in scenarios like callbacks or functional programming tasks.

## Contributing

Would you like to contribute to this project? Great! We warmly welcome your contributions. Follow the steps below to get started:  

1. **Fork the repository**  
   Fork the project to your GitHub account.  

2. **Clone the repository**  
   Clone the forked repository to your local machine:  
   ```bash
   git clone https://github.com/username/project-name.git

3. **Create a branch**
    Create a branch for your feature or fix:
    ```bash
    git checkout -b feature-name
    ```
4. **Make your changes**
    ```bash
    git add .
    git commit -m "Descriptive commit message"
    ```
5. **Push your branch to the remote repository**
    ```bash
    git push origin feature-name
    ```
6. **Create a Pull Request**\
    Open the repository on GitHub and create a "New Pull Request." Include a clear description of the changes you’ve made.

# Resources
- https://go.dev
- https://go.dev/dl/
- https://www.youtube.com/watch?v=AD2zUrhL6GI