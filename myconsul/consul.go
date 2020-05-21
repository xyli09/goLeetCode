package myconsul


import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
	"log"
	"net"
	"strconv"
	//"testing"  t *testing.T
)

const Id = "1234567890"

func Register() {

	fmt.Println("test begin .")
	config := consulapi.DefaultConfig()
	//config.Address = "localhost"
	fmt.Println("defautl config : ", config)
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal("myconsul client error : ", err)
	}
	//创建一个新服务。
	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = Id
	registration.Name = "user-tomcat"
	registration.Port = 8080
	registration.Tags = []string{"user-tomcat"}
	registration.Address = "127.0.0.1"

	////增加check。
	//check := new(consulapi.AgentServiceCheck)
	//check.HTTP = fmt.Sprintf("http://%s:%d%s", registration.Address, registration.Port, "/check")
	////设置超时 5s。
	//check.Timeout = "5s"
	////设置间隔 5s。
	//check.Interval = "5s"
	////注册check服务。
	//registration.Check = check
	//log.Println("get check.HTTP:",check)

	err = client.Agent().ServiceRegister(registration)



	if err != nil {
		log.Fatal("register server error : ", err)
	}

}

func Dregister(){


	fmt.Println("test begin .")
	config := consulapi.DefaultConfig()
	//config.Address = "localhost"
	fmt.Println("defautl config : ", config)
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal("myconsul client error : ", err)
	}

	err = client.Agent().ServiceDeregister("test001")
	if err != nil {
		log.Fatal("register server error : ", err)
	}


}



func ReadConsul() {
	var lastIndex uint64
	config := consulapi.DefaultConfig()
	config.Address = "127.0.0.1:8500" //myconsul server

	client, err := consulapi.NewClient(config)
	if err != nil {
		fmt.Println("api new client is failed, err:", err)
		return
	}
	services, metainfo, err := client.Health().Service("test", "dev", true, &consulapi.QueryOptions{
		WaitIndex: lastIndex, // 同步点，这个调用将一直阻塞，直到有新的更新
	})
	if err != nil {
		logrus.Warn("error retrieving instances from Consul: ", err)
	}
	lastIndex = metainfo.LastIndex

	addrs := map[string]struct{}{}
	for _, service := range services {
		fmt.Println("service.Service.Address:", service.Service.Address, "service.Service.Port:", service.Service.Port)
		addrs[net.JoinHostPort(service.Service.Address, strconv.Itoa(service.Service.Port))] = struct{}{}
	}
}