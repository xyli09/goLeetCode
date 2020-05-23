package HWtest

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type info struct {
	fullName string
	name string
	line int
	count int
}

func SecondQuestion() {

	var (
		N ,M string
		Input string
		//Names  []string

	)
	infoMap := make(map[string]info)

	f := bufio.NewReader(os.Stdin) //读取输入的内容

	for {
		//fmt.Scanf("%s %s",&N,&M)//临时存储
		Input,_ = f.ReadString('\n') //定义一行输入的内容分隔符。
		if len(Input) == 0 {
			break //如果用户输入的是一个空行就让用户继续输入。
		}
		fmt.Sscan(Input,&N,&M) //将Input
		fname := getFullName(N)
		name := getName(fname)
		line,_ := strconv.Atoi(M)
		key := fname+strconv.Itoa(line)
		if v,ok := infoMap[key];ok{

			ins := v
			ins.count = v.count+1
			infoMap[key] = ins
		}else{
			ins := info{
				fullName: fname,
				name:     name,
				line:     line,
				count:    1,
			}
			infoMap[key] = ins
		}

	}

	for k,_ := range(infoMap){
		printRes(infoMap[k])
	}



}

func printRes(ins info){
	fmt.Println(ins.name+" "+strconv.Itoa(ins.line)+" "+strconv.Itoa(ins.count))
}

func getFullName(fname string)string{
	ss := strings.Split(fname,"\\")
	return ss[len(ss)-1]
}

func getName(fname string)string{
	ss := []rune(fname)
	ssl := len(ss)
	if ssl <= 16{
		return fname
	}else{
		return string(ss[ssl-1:ssl])
	}
}