package main

import (
	"fmt"

	"github.com/Windmill787-golang/bubble-sort/bubble"
)

func main() {
	s := []int{1, 2, 3, 15, 7, 2, 1, 8, 9}

	fmt.Println(bubble.Sort(s))
}
