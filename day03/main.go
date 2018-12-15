package main

import (
  "io/ioutil"
  "fmt"
  //"math"
  //"strings"
  //"strconv"
  //"sort"
  )

func main() {
	north := "^"
	south := "v"
	east  := ">"
	west := "<"

	santaX := 128
	santaY := santaX
	roboX := santaX
	roboY := santaY

	size := santaX*2

    grid := make([][]int, size)
    for i,_ := range grid {
    	grid[i] = make([]int, size)
    	for j,_ := range grid[i] {
    		grid[i][j] = 0
    	}
    }

    input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\r\n",err)
	} else {
		// Deliver initial
		grid[santaX][santaY] = 2;

		for i,c := range input {

			x := santaX
			y := santaY

			if i%2 == 0 {
				x = roboX
				y = roboY
			}

			char := string(c)
			if char == north {
				y++
			} else if char == south {
				y--
			} else if char == east {
				x++
			} else if char == west {
				x--
			} else {
				fmt.Print("nope......................")
			}
			// Deliver next
		    grid[x][y] = grid[x][y] + 1;

		    if i%2 == 0 {
				roboX = x
				roboY = y
			} else {
				santaX = x
				santaY = y
			}
		}
	}

 	luckyKids := 0
 	
    for i,_ := range grid {
    	for j,_ := range grid[i] {
    		house := grid[i][j]
    		if house > 0 {
    			luckyKids++
    		}
    		fmt.Printf("%v",house)
    	}
    	fmt.Printf("\r\n")
    }	

    fmt.Printf("Lucky kids: %v\r\n",luckyKids)
}
