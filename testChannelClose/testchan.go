package testChannelClose


import (
	"fmt"
	"strconv"
	"time"
)

var ch chan int

func f() {
	ch = make(chan int, 5)
	println("start")
	//for{
		for i := 1; i < 5; i++ {
			ch <- i
		}
		time.Sleep(time.Second * 10)
	//}

	close(ch)
	println("close")
}

func isClosed(ch1 chan int)  {
	select {
	case <-ch1:
		fmt.Println("is open")
	default:
		fmt.Println("is closed")
	}

}


func TestChanClose() {
	go f()

	time.Sleep(time.Second * 1)
	println("closed")

	isClosed(ch)


	for c := range ch {
		println(c)
	}
}


func TestChan2() {
	go f()
	//close(c)
	//c <- 5
	time.Sleep(time.Second * 1)
	n := []int{1,2,3,4}
	for k := range(n){
		fmt.Println("this is line:"+strconv.Itoa(k+1))
		//isClosed(ch)
		i, ok := <-ch
		if ok {
			fmt.Println("channel is ok!")
			fmt.Println("ok value:"+strconv.Itoa(i))
		}else{
			fmt.Println("channel closed!")
			fmt.Println("chan value:"+strconv.Itoa(i))
		}

	}


	isClosed(ch)

	i, ok := <-ch
	if ok {
		fmt.Println("channel is ok!")
	}else{
		fmt.Println("channel closed!")
		fmt.Println("last")
		fmt.Println(i)
	}

	a := <- ch
	fmt.Println(a)
}
