package numberfunctions

import (
	"fmt"
	"math/rand"
	"time"
)

func CreateListThousandLargeNumbers() string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	var numList [1000]int64
	var strNumList string

	for i := 0; i < 1000; i++ {
		// generate random number in the range 1e14 (100 trillion) to 1e15-1 (999 trillion)
		numList[i] = r.Int63n(9e14) + 1e14
		strNumList += fmt.Sprintf("%d ", numList[i])
	}

	return strNumList

}
