package HWtest

import (
	"fmt"
	"strconv"
)

func FirstQuestion() {
	var (
		N ,M int
		H []int //score
		res []int
		//Names  []string

	)
	fmt.Scanln(&N,&M)

	for i := 0; i < N; i++ {
		var score string

		fmt.Scanf("%s",&score)//临时存储
		h,_ := strconv.Atoi(score)
		H = append(H, h)//追加到切片中
	}
	//fmt.Println(H)


	for i := 0; i < M; i++ {
		var act string
		var s, e int
		fmt.Scanln(&act, &s, &e)//临时存储
		if act == "Q"{
			if s > e{
				s,e = e,s
			}
			m := getMax(H,s,e)
			res = append(res,m)
		}
		if act == "U"{
			H[s-1] = e
		}

	}

	for _,v := range res{
		fmt.Println(v)
	}


}



func getMax(table []int,start,end int)int{
	if table == nil{
		return 0
	}
	max := 0
	if end > len(table){
		return 0
	}
	for i:= start-1;i<end;i++{
		if max < table[i]{
			max = table[i]
		}
	}
	return max
}