package main

import (
	"fmt"
	"regexp"
)

func main() {
	a := "I am learning Go language"

	re, _ := regexp.Compile("[a-z]{2,4}")

	// Find the first match
	one := re.Find([]byte(a))
	fmt.Println("Find:", string(one))
	// Find all matches and save them in a segment, n less than 0 means all matches, indicate the segment size if it is greater than 0.
	all := re.FindAll([]byte(a), -1)
	fmt.Println("FindAll", all)

	// Finds the index of the first match, start and end position.
	index := re.FindIndex([]byte(a))
	fmt.Println("FindIndex", index)

	// Finding the indices of all matches, n does the same as above
	allindex := re.FindAllIndex([]byte(a), -1)
	fmt.Println("FindAllIndex", allindex)

	re2, _ := regexp.Compile("am(.*)lang(.*)")

	// Finds the first submatch and returns an array, the first element contains all the elements, the second contains the result of the first (), and the third contains the results of the second ()
	// Output:
	// First element: "am learning Go lenguage"
	// Second element: "learning Go"
	// Third element: "usage"
	submatch := re2.FindSubmatch([]byte(a))
	fmt.Println("FindSubmatch", submatch)
	for _, v := range submatch {
		fmt.Println(string(v))
	}

	// Same as FindIndex().
	submatchindex := re2.FindSubmatchIndex([]byte(a))
	fmt.Println(submatchindex)

	// FindAllSubmatch
	submatchall := re2.FindAllSubmatch([]byte(a), -1)
	fmt.Println(submatchall)

	// FindAllSubmatchIndex
	submatchallindex := re2.FindAllSubmatchIndex([]byte(a), -1)
	fmt.Println(submatchallindex)
}
