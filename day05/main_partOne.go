package main

import (
  "io/ioutil"
  "fmt"
  "strings"
  "regexp"
  )

func hasThreeVowels(input string) bool {
	vowels := "aeiou"
	vCount := 0
	for _,c := range input {
	   for _,v := range vowels {
	   		if c == v {
	   			vCount++
	   		}
	   }
	}
	return vCount > 2
}

func hasDoubles(input string) bool {
	prev := '@'
	for i,c := range input {
		if i > 0 {
			if c == prev {
				return true
			}
		}
		prev = c
	}
	return false
}

func containsNaughty(input string) bool {
	pattern := "ab|cd|pq|xy"
	matches, _ := regexp.MatchString(pattern, input)
	return matches
}

func isNice(input string) bool {
	result := hasThreeVowels(input)
	if result {
		result = hasDoubles(input)
	}
	if result {
		result = !containsNaughty(input)
	}
	return result
}

func main() {

    nice := 0
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error: %v\r\n",err)
	} else {
		
		lines := strings.Split(string(input),"\r\n")



		for _,line := range lines {
			if isNice(line) {
				nice++
				fmt.Printf("NICE:%v\r\n",line)
			}
	    }
	}

	fmt.Printf("\r\n All nice: %v",nice)
}

