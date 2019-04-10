package usefulFunc

import "fmt"

func ShowSlice(sli []int) {
	for _, v := range sli {
		fmt.Println(v)
	}
}
