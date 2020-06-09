package testSlice

import (
	"encoding/json"
	"fmt"
)

func TestSliceNil(){
	var s []int
	ls := len(s) // 0
	cs := cap(s) // 0
	fmt.Println(ls,cs)
	for range s {
	} // iterates zero times
	//s[i] // panic: index out of range
}

func TestSliceAppend(){
	var s []int
	s = nil
	sstr,_ := json.Marshal(s)  //null
	fmt.Println("s1nil:",string(sstr))

	for i := 0; i < 10; i++ {
		fmt.Printf("len: %2d cap: %2d\n", len(s), cap(s))
		s = append(s, i)
	}
	sstr,_ = json.Marshal(s)
	fmt.Println(string(sstr))


	s2 := make([]int,0,0)
	s2str,_ := json.Marshal(s2)  //  []
	fmt.Println("s2nil:",string(s2str))

	for i := 0; i < 10; i++ {
		fmt.Printf("len: %2d cap: %2d\n", len(s2), cap(s2))
		s2 = append(s2, i)
	}

	s2str,_ = json.Marshal(s2)
	fmt.Println(string(s2str))
}
