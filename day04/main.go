package main

import (
  "fmt"
  "crypto/md5"
  "encoding/hex"
  )

// Credit to sergiotapia for this little nugget that saved me some time.
// https://gist.github.com/sergiotapia/8263278

func main() {
	input := "bgvyzdsv"
	i := 0

	for {
		text := input + fmt.Sprintf("%05d",i)
//		fmt.Println(text)
		hasher := md5.New()
		hasher.Write([]byte(text))

		outString := hex.EncodeToString(hasher.Sum(nil))
		subString := outString[0:6]
		if subString == "000000" {
			fmt.Println("Hooray!")
			break
		} else {
			i++
		}
	}

	fmt.Printf("lowest input is : %d\r\n",i)
}
