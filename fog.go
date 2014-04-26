package main

import "fmt"

func main() {
	lookingFor := uint64(910897038977002)
	letters := "acdegilmnoprstuww"
	
	fmt.Println("IterativeMethod(): ")
	iterativeMethod(lookingFor, letters)
	fmt.Println()
	
	fmt.Println("logarithmicMethod(): ")
	logarithmicMethod(lookingFor, letters)
}

/*
	Since the hashing function only has 16 input values per position and 
	multiplies the value by 37, there will be no overlap in hashing results 
	for any combination of inputs.
*/
func iterativeMethod(lookingFor uint64, letters string) {
	first := letters[0]
	last := letters[len(letters) - 1]
	threshold := generate(rune(first), len(letters)) + string(last)
	sureAbout := ""
	for charsFound := 0; charsFound < 9; charsFound++ {
		var lastChar rune
		for j, c := range letters {
			trial := sureAbout + string(c) + generate(rune(threshold[j]), 8 - charsFound)
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
	first := letters[0]
	last := letters[len(letters) - 1]
	threshold := generate(rune(first), len(letters)) + string(last)
	sureAbout := ""
	result := logarithmicMethodRec(lookingFor, letters, threshold, 0, len(letters), sureAbout)
	fmt.Println(result)
}

func logarithmicMethodRec(lookingFor uint64, letters string, threshold string, startAt int, stopAt int, sureAbout string) string {
	charsFound := len(sureAbout)
	midpoint := (stopAt - startAt) / 2 + startAt
	candidate := string(letters[midpoint])
	filler := rune(threshold[midpoint])
	trial := sureAbout + candidate + generate(filler, 8 - charsFound)
	result := hash(trial)

	fmt.Printf("Trying '%s'\n", trial)
	if result == lookingFor {
		return trial
	} else if midpoint == startAt {
		sureAbout = sureAbout + candidate
		return logarithmicMethodRec(lookingFor, letters, threshold, 0, len(letters), sureAbout)
	} else if result > lookingFor {
		return logarithmicMethodRec(lookingFor, letters, threshold, startAt, midpoint, sureAbout)
	} 
	
	return logarithmicMethodRec(lookingFor, letters, threshold, midpoint, stopAt, sureAbout)
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