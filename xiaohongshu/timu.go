package xiaohongshu



func Find(a,b []int){
	la := len(a)
	lb := len(b)
	pa,pb := 0,0
	res := make([]int,0,len(a))
	for  pa<la&&pb<lb{

		if a[pa] == b[pb]{
			pb += 1
			pa += 1
			continue
		}
		if a[pa] < b[pb]{
			res = append(res, a[pa])
			pa += 1
			continue
		}else{
			pb+=1
		}
	}

	if pa < la{
		for i:= pa ;i<la;i++{
			res = append(res, a[i])
		}
	}
	//res = append(res, a[pa:])
	for _,n := range(res){
		println(n)
	}
}

func TestFind(){
	a := []int{1,2,4,5,8,10}
	b := []int{2,3,8,9}
	Find(a,b)
}
