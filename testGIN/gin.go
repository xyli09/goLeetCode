package testGIN

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
)

func Request(){
	request := gorequest.New()
	resp, body, err := request.Get("http://192.168.0.44:8080/v1/getinfo").

		Set("If-None-Match", `W/"wyzzy"`).
		End()
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(resp)
	fmt.Println(body)
}
