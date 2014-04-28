package main

import "fmt"

func main() {
	lookingFor := uint64(910897038977002)
	letters := "acdegilmnoprstuww"
		
	fmt.Println("iterativeMethod(): ")
	iterativeMethod(lookingFor, letters)
	
	fmt.Println("logarithmicMethod(): ")
	logarithmicMethod(lookingFor, letters)

	fmt.Println("reverseMethod(): ")
	reverseMethod(lookingFor, letters)

	fmt.Println("bonus: ")
	bonushash := uint64(34700290059707989)
	reverseMethod(bonushash, letters)
}

/*
	Since the hashing function only has 16 input values per position and 
	multiplies the value by 37, there will be no overlap in hashing results 
	for any combination of inputs.
*/
func iterativeMethod(lookingFor uint64, letters string) {
	wordlength := determineWordlength(lookingFor, letters)
	first := letters[0]
	last := letters[len(letters) - 1]
	threshold := generate(rune(first), len(letters)) + string(last)
	sureAbout := ""
	for charsFound := 0; charsFound < wordlength; charsFound++ {
		var lastChar rune
		for j, c := range letters {
			trial := sureAbout + string(c) + generate(rune(threshold[j]), wordlength - charsFound - 1)
			fmt.Printf("Trying '%s': ", trial)
			result := hash(trial)
			if result == lookingFor {
				fmt.Println("PASS!")
				return
			} else if result > lookingFor {
				sureAbout = sureAbout + string(lastChar)
				fmt.Printf("FAIL but found letter '%c'!\n", lastChar)
				break
			} else {
				lastChar = c
			}
			fmt.Println("FAIL!")
		}
	}
}

func logarithmicMethod(lookingFor uint64, letters string) {
	wordlength := determineWordlength(lookingFor, letters)
	first := letters[0]
	last := letters[len(letters) - 1]
	threshold := generate(rune(first), len(letters)) + string(last)
	sureAbout := ""
	result := logarithmicMethodRec(lookingFor, letters, wordlength, threshold, 0, len(letters), sureAbout)
	fmt.Println(result)
}

func logarithmicMethodRec(lookingFor uint64, letters string, wordlength int, threshold string, startAt int, stopAt int, sureAbout string) string {
	charsFound := len(sureAbout)
	midpoint := (stopAt - startAt) / 2 + startAt
	candidate := string(letters[midpoint])
	filler := rune(threshold[midpoint])
	trial := sureAbout + candidate + generate(filler, wordlength - charsFound - 1)
	result := hash(trial)

	fmt.Printf("Trying '%s'\n", trial)
	if result == lookingFor {
		return trial
	} else if midpoint == startAt {
		sureAbout = sureAbout + candidate
		return logarithmicMethodRec(lookingFor, letters, wordlength, threshold, 0, len(letters), sureAbout)
	} else if result > lookingFor {
		return logarithmicMethodRec(lookingFor, letters, wordlength, threshold, startAt, midpoint, sureAbout)
	} 
	
	return logarithmicMethodRec(lookingFor, letters, wordlength, threshold, midpoint, stopAt, sureAbout)
}

func determineWordlength(lookingFor uint64, letters string) int {
	first := letters[0]
	for i := 1; i < 12; i++ {
		candidate := generate(rune(first), i)
		h := hash(candidate)
		if lookingFor < h {
			return i - 1
		}
	} 
	return 0
}

func reverseMethod(lookingFor uint64, letters string) {
	result := reverseMethodRec(lookingFor, letters)
	if result == "" {
		fmt.Println("No result found!")
	} else {
		fmt.Println(result)
	}
}

func reverseMethodRec(lookingFor uint64, letters string) string {
	if lookingFor < 8 {
		return ""
	}
	for i, c := range letters {
		trial := lookingFor - uint64(i)
		if trial % 37 == 0 {
			return reverseMethodRec(trial / 37, letters) + string(c)
		}
	}
	return ""
}

// As per provided pseudo-code
func hash(s string) uint64 {
	var h uint64 = 7
	
	letters := "acdegilmnoprstuw"
	for _, r := range s {
		h = (h * 37 + indexOf(letters, r))
    }

	return h
}

// Ignore this code below, as I have no idea how to do this in go
// Equivalent to new String(char, count) in c#
func generate(c rune, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result = result + string(c)
	}
	return result
}

// No explanation needed...
// Although this could be optimized with a lookup table
func indexOf(s string, c rune) uint64 {
	for i, r := range s {
		if r == c {
			return uint64(i)
		}
	}
	return 0
}