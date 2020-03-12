package main

import (
	"fmt"
	"webserver/morestrings"

	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("okO"))
	fmt.Println(cmp.Diff("HelloWorld", "Sleep"))
}
