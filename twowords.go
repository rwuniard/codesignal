package main

import "fmt"

func main() {
	fmt.Println(solution("aabg", "cdef"))

	fmt.Println(solution("abg", "ccdeff"))
	fmt.Println(solution("abg", "cddeff"))
}

func solution(s1 string, s2 string) string {
	result := ""
	i := 0
	j := 0
	for i, j = 0, 0; i < len(s1) && j < len(s2); {
		// If s1 doesn't have a repeat character, then s2 go to the result.
		if checkRepeat(s1, i) {
			result += string(s2[j:])
			j = len(s2)
		} else if checkRepeat(s2, j) { // If s2 doesn't have a repeat character, then s1 go to the result.
			result += string(s1[i:])
			i = len(s1)
		} else if s1[i] < s2[j] { // If s1 and s2 doesn't have a repeat character then whichever char is smaller go to the result.
			result += string(s1[i])
			i++
		} else {
			result += string(s2[j])
			j++
		}

		// If one of the string has been exhausted, then the other string go to the result.
		if i == len(s1) {
			result += s2[j:]
		} else if j == len(s2) {
			result += s1[i:]
		}
		fmt.Println("i:", i, "j:", j)
	}
	return result
}

func checkRepeat(s string, pos int) bool {
	// Check if there is a repeat character in s.
	// If there is a repeat character, return true.
	if pos == len(s)-1 {
		return false
	}
	if s[pos] == s[pos+1] {
		return true
	}
	return false
}
