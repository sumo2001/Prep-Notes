## Introduction
- Course's XML -> https://github.com/StephenGrider/GoCasts/tree/master/diagrams
```
    package main
    import "fmt"
    func main(){
     fmt.Println("This is the output boi")
    }
 ```
 - The package = project = workspace,there are 2 types of packages 1. spits an executable 2. reusbale packages the main purpose of package is to group together code with similar process
 - main is a special type of package which spits out a special type of executable
 - ![image](https://user-images.githubusercontent.com/51809378/172064045-368091ba-826c-4f1a-9ba5-e33a4cffc438.png)
 - FMT is a type of depedency which gets imported upon call, there are many dependencies like calc, os, math etc..
 - the func main is the important thing, everything we write under it gets executed
 - ![image](https://user-images.githubusercontent.com/51809378/172064132-6c944b0c-bfa5-4252-a65a-cd6514446d1d.png)
 ## Digging Deeper
 - Go is a static language which is different from dynamic languages like python, java, js, in go you need to tell which data type the variable is going to save
 ```
 In js: var number = "abcd"
 In go: var card string = "abcd" or card := "abcd"
 ```
 - Go can automatically reconginze the data type from the input of right side and automatically assigns the variable type to it and also := is only used to assign a new variable
 - In Go, the functions within same packages ad different files can call each other, without defining them in the each others code
 - Array are not editable, where slices are a from of arrays with ediatble features
 ``` 
  - Create a slice
 cards := []string{"data 1", "data 2"}
  -  Append data
  cards = append(cards, "new data")
  - For Loop for a slice
  cards := []string{"data 1", "data 2"}
  for i, card := range cards {
    printLn{i, card}
  } 
 ``` 
 - Compiler will yell at you if you dont use the variable that you declared in your code

