package maxSubString

import (
	"fmt"
	"strings"
)

func LengthOfLongestSubString(s string) int {
	var Length int
	var s1 string
	left := 0
	right := 0
	s1 = s[left:right]
	for ; right < len(s); right++ {
		ts :=  s[right]
		tstr := string(ts)
		fmt.Println(tstr)
		index := strings.IndexByte(s1,ts)
		if  index != -1 {
			left += index + 1
		}
		s1 = s[left : right+1]
		if len(s1) > Length {
			Length = len(s1)
		}
	}

	return Length

}


func TestLongestSubString() {
	s := "ejaajeq"

	maxl := LengthOfLongestSubString(s)
	fmt.Println(maxl)
}