package followup

import "fmt"

func follow(data []int){
	slow :=0
	n := len(data)
	for fast := 0; fast<n;fast++{
		if data[fast]!=0{
			data[slow] = data[fast]
			slow++
		}
	}
	for ;slow<n;slow++{
		data[slow]=0
	}
	fmt.Println(data)
}

func TestFollow(){
	data := []int{1, 0, 5, 0, 13,17}
	follow(data)
}


