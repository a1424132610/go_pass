package main

import (
	context "context"
	"fmt"
	"github.com/a1424132610/go_common/config"
	"github.com/a1424132610/go_pass/pod/proto/pod"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/registry"
	"github.com/go-micro/plugins/v3/registry/consul"
)

type Pod struct{}

func (p *Pod) SayHello(ctx context.Context, request *pod.HelloRequest, response *pod.HelloResponse) error {
	return nil
}

func main() {
	// 创建 Consul 注册中心
	consulRegistry := consul.NewRegistry(
		registry.Addrs("192.168.254.130:8500"), // 只设置可访问的节点
	)
	// 创建新的服务
	service := micro.NewService(
		micro.Name("example"),
		micro.Registry(consulRegistry),
	)
	config, err := config.NewConsulConfig(
		"192.168.254.130", 8500, "/micro/config",
	).GetConfig()
	if err != nil {
		logger.Error(err)
	}
	fmt.Println(config)
	pod.RegisterGreeterHandler(service.Server(), new(Pod))
	// 初始化服务
	service.Init()
	err = service.Run()
	if err != nil {
		logger.Error(err)
	}
}
