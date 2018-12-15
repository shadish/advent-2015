package main

import (
  "io/ioutil"
  "fmt"
  "strings"
//  "regexp"
  )

func sandwich(input string) bool {
	return false
}

func hasDoubles(input string) bool {
  //fmt.Printf("%v :: ",input)

  sets := []string{}

  l := len(input)
	for i := 0; i < l; i++ {
    if i < l-1 {
      sub := input[i:i+2]
      sets = append(sets,sub)
	  	//fmt.Printf("%v|",sub)
    }
	}

  for _,se := range sets {
    found := 0
    for i := 0; i < l; i++ {
      if i < l-1 {
        sub := input[i:i+2]
        if sub == se {
          found++
        }
      }
    }
    if found > 1 {
      fmt.Printf("(%v)", se)
      return true
    }
  }
  
  //fmt.Printf("\r\n")
  return false
}

func isNice(input string) bool {
	result := hasDoubles(input)
	if result {
		//result = sandwich(input)
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
