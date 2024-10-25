
### Introduction To Backend With Golang

\- Anurag Rao

___
### Stuff We're Gonna Do Today

1. Golang Syntaxes and Nuances
	- Packages
	- Variables & Types
	- Short Hand For Variable Declaration
	- Conditionals (If Statements)
	- Loops
	- Arrays
1. Create a program that takes user input for username & password and store them in an array. Then display them at the end. 

___

### Golang Syntaxes And Nuances

___

###### Packages

- Packages are names that a bunch of files share
- Functions and variables are shared between these files
- A powerful feature in golang not present in other languages
- Basically, just defines _identifier visibility?_

___

```go
// file1.go
package foo

func bar() {
	fmt.Println("Hello!")
}
```

```go
//file2.go

package foo

bar() // I can call the function defined in the other file
```

- Helps with code organization

___
