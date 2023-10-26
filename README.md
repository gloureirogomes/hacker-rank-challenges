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

To run a challenge, in the 'main.go' file at the root of the project, uncomment the desired line and import 

```golang
package main

import (
	"fmt"
	// birthdaycakecandles "hacker-rank/birthday_cake_candles"
	// comparethetriplets "hacker-rank/compare_the_triplets"
	// diagonaldifference "hacker-rank/diagonal_difference"
	// simplearraysum "hacker-rank/simple_array_sum"
	solvemefirst "hacker-rank/solve_me_first"
	// timeconversion "hacker-rank/time_conversion"
	// verybigsum "hacker-rank/very_big_sum"
)

func main() {
	fmt.Println("Hacker Rank Challenges!!!")

	// Uncomment the line below and the import to run the desired package and function
	fmt.Println(solvemefirst.SumTwoNumbers())
	// fmt.Println(simplearraysum.SumArray())
	// fmt.Println(comparethetriplets.CompareTriplets())
	// fmt.Println(verybigsum.VeryBigSum())
	// fmt.Println(diagonaldifference.DiagonalDifference())
	// plusminus.PlusMinus()
	// staircase.Staircase()
	// minimaxsum.MiniMaxSum()
	// fmt.Println(birthdaycakecandles.BirthdayCakeCandles())
	// fmt.Println(timeconversion.TimeConversion())
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
5
```
 Note: The entries used in the challenges are the entries suggested by the challenge itself in HackerRank 