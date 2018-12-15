package main

import (
  "io/ioutil"
  "fmt"
  "math"
  "strings"
  "strconv"
  "sort"
  )

func surface(l int,w int,h int) int {
	lSide := float64(l*w)
	wSide := float64(w*h)
	hSide := float64(h*l)
	baseArea := (lSide + wSide + hSide)*2

	slack := math.Min(lSide, wSide)
	slackA := math.Min(slack, hSide)

	baseArea += slackA
	return int(baseArea)
}


func ribbon(dimensions []int) int {
	sort.Ints(dimensions)
	sorted := dimensions
	result := sorted[0]*2 + sorted[1]*2
	cubic := sorted[0]*sorted[1]*sorted[2]
	fmt.Printf("ITS %v + %v ... %v\r\n",sorted[0],sorted[1],result)
	return result+cubic
}

func main() {
    total := 0
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error: %v\r\n",err)
	} else {
		
		lines := strings.Split(string(input),"\r\n")
		for _,line := range lines {
			values := strings.Split(line,"x")
			l,_ := strconv.ParseInt(values[0],10,64)
			w,_ := strconv.ParseInt(values[1],10,64)
			h,_ := strconv.ParseInt(values[2],10,64)

			vals := []int {int(l),int(w),int(h)}
			rib := ribbon(vals)
			//surf := surface(int(l),int(w),int(h))
			total += rib
			fmt.Printf("Found value: %v x %v x %v ... %v :: Total: %v\r\n",l,w,h,rib,total)
	    }
	}

	//vals := []int {1,1,10}
    //rib := ribbon(vals)
    //fmt.Printf("so.... %v",rib)

	fmt.Printf("total: %v\r\n",total)
}
