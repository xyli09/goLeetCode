package main

import (
	"fmt"
	"strconv"
)

//func main(){
//
//	//s := []int{2, 3, 5, 7, 11, 13}
//	//testGIN.Request()
//	//testMapStruct.Map2Struct()
//	s :="ejaajeq"
//	l:=huiwen.LongestPalindromeSubseq(s)
//	fmt.Println(l)
//
//	maxl := maxSubString.LengthOfLongestSubstring(s)
//	fmt.Println(maxl)
//
//	//myconsul.ReadConsul()
//
//	newecc.MyECC()
//}

//func main() {
//	ctx, cancel := context.WithCancel(context.Background())
//	go func(ctx context.Context) {
//		for {
//			select {
//			case <-ctx.Done():
//				fmt.Println("监控退出，停止了...")
//				return
//			default:
//				fmt.Println("goroutine监控中...")
//				time.Sleep(2 * time.Second)
//			}
//		}
//	}(ctx)
//
//	time.Sleep(10 * time.Second)
//	fmt.Println("可以了，通知监控停止")
//	cancel()
//	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
//	time.Sleep(5 * time.Second)
//
//}

func main() {
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