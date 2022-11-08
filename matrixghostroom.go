package main

import "fmt"

/*
After becoming famous, the CodeBots decided to move into a new building together.
Each of the rooms has a different cost, and some of them are free, but there's a
rumour that all the free rooms are haunted! Since the CodeBots are quite superstitious,
they refuse to stay in any of the free rooms, or any of the rooms below any of the free rooms.

Given matrix, a rectangular matrix of integers, where each value represents the cost of the room,
your task is to return the total sum of all rooms that are suitable for the CodeBots (ie: add up all
the values that don't appear below a 0).

*/

func main() {
	test1 := [][]int{{0, 1, 1, 2}, {0, 5, 0, 0}, {2, 0, 3, 3}}
	test2 := [][]int{{1, 1, 1, 0}, {0, 5, 0, 1}, {2, 1, 3, 10}}
	fmt.Println(matrixElementsSum(test1))
	fmt.Println(matrixElementsSum(test2))
}

func matrixElementsSum(matrix [][]int) int {
	// your code here
	// create a map of the haunted rooms
	hauntedRooms := make(map[int]bool)
	// create a variable to hold the sum
	sum := 0
	// iterate through the matrix
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			// if the room is haunted, then add it to the map
			if matrix[i][j] == 0 {
				hauntedRooms[j] = true
				fmt.Println("i, j", i, j)
			}
			// if the room is not haunted, then add it to the sum
			if _, ok := hauntedRooms[j]; !ok {
				fmt.Println("not haunted", i, j, matrix[i][j])
				sum += matrix[i][j]
			}
		}
	}
	return sum
}
