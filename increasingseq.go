package main

import "fmt"

func main() {
	test1 := []int{3, 5, 67, 98, 3}
	test2 := []int{1, 2, 3, 4, 3, 6}
	fmt.Println(solution(test1))
	fmt.Println(solution(test2))
}

// function for strictly increasing sequence from Code Signal
// https://app.codesignal.com/arcade/intro/level-2/2mxbGwLzvkTCKAJMG
func solution(sequence []int) bool {
	bad := 0
	for i := 1; i < len(sequence); i++ {
		if sequence[i] <= sequence[i-1] {
			bad++
		}
		if bad > 1 {
			return false
		}
		if (i-1 >= 0) && (i+2 <= len(sequence)-1) && (sequence[i]-sequence[i+2] >= 0) && (sequence[i-1]-sequence[i+1] >= 0) {
			return false
		}
	}
	return true
}

// // This is the logic for strictly increasing sequence from Code Signal
// func main() {

// 	// check if the array is a strictly increasing sequence
// 	// if it is, then return true
// 	// if it is not, then return false
// 	test1 := []int{3, 5, 67, 98, 3}
// 	test2 := []int{1, 2, 3, 4, 3, 6}
// 	fmt.Println(solution(test1))
// 	fmt.Println(solution(test2))
// }
// func solution(sequence []int) bool {
// 	strictSq, pos := isStrictlyIncreasingSequence(sequence)
// 	if !strictSq {
// 		for i := pos; i < len(sequence); i++ {
// 			fmt.Println("original sequence", sequence)
// 			newSq := RemoveIndex(sequence, i)
// 			fmt.Println("New:", newSq, i)
// 			fmt.Println("Old:", sequence)
// 			newStrictSq, _ := isStrictlyIncreasingSequence(newSq)
// 			if newStrictSq {
// 				return true
// 			}
// 		}
// 	}
// 	return strictSq
// }

// func RemoveIndex(s []int, index int) []int {
// 	myCopy := make([]int, len(s))
// 	copy(myCopy, s)
// 	return append(myCopy[:index], myCopy[index+1:]...)
// }

// func isStrictlyIncreasingSequence(arr []int) (bool, int) {
// 	// your code here
// 	if len(arr) == 0 || len(arr) == 2 {
// 		return true, -1
// 	}
// 	for i := 0; i < len(arr)-1; i++ {
// 		if arr[i] >= arr[i+1] {
// 			return false, i
// 		}
// 	}
// 	return true, -1
// }
