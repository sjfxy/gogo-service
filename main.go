package main

import (
	"fmt"
	"os"

	"github.com/cloudfoundry-community/go-cfenv"
	service "github.com/sjfxy/gogo-service/service"
)

//2222
func main() {
	//调用系统的PORT对应的环境变量 我们可以在构建云的时候可以使用对应的外部的环境变量处理对应的分配的端口号和负载均衡的策略
	//如果没有的话则默认在云端会启动3000端口的
	//使用的是cloundfoundry-comou社区的go-cfenv的处理方式
	//可以使用本地进行推送即可
	//我们使用 cf的相关处理方式进行部署的反射
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	appEnv, err := cfenv.Current()
	if err != nil {
		fmt.Println("CF Environment not detected.")
	}

	server := service.NewServer(appEnv)
	server.Run(":" + port)
}
