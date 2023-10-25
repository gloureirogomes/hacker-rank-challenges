# Hacker-Rank-Challenges

Project to save the implementations of the https://www.hackerrank.com challenges made in Golang.

## Installation

Make sure you have Golang version **^1.21.3** installed on your machine.

After cloning the project, go to the folder and run the following commands so that Go recognizes the project and downloads the necessary dependencies

```bash
  cd hacker-rank-challenges
  go mod download
  go mod tidy
```

## Stack used

**Back-end:** Golang

## Use/Examples

To run a challenge, at the root of the project, in the 'main.go' file, import the desired package and call the only function in the package.

Note: Remember to add the necessary inputs for the function to perform correctly and also some functions don't print the final answer, so make sure to add the function `fmt.Println()`

```golang
package main

import (
	"fmt"
	solvemefirst "hacker-rank/solve_me_first"
)

func main() {
	fmt.Println("Hacker Rank Challenges!!!")

	inputA := uint32(5)
	inputB := uint32(10)
	fmt.Println(solvemefirst.SumTwoNumbers(inputA, inputB))
}
```

Once this is done, just run the following command in the terminal at the root of the project:

```bash
  go run main.go
```

**Output**

```bash
❯ go run main.go
Hacker Rank Challenges!!!
15
```
