package main

import (
	"context"
	"fmt"
	"github.com/ameamezhou/xiawuyue"
	"github.com/ameamezhou/xiawuyue/xlog"
	"github.com/careerplan/comm"
	"github.com/careerplan/model"
	"github.com/careerplan/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main(){
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s [conf]\n", os.Args[0])
		return
	}

	conf := os.Args[1]
	svr := xiawuyue.New()
	// 读配置
	e := comm.LoadSettings(conf)
	if e != nil {
		fmt.Printf("Parse config fail: %v\n", e)
		return
	}
	// 环境变量初始化
	svr.EnvInit(comm.Config)
	model.Init()
	defer model.Close()

	// 进程上下文管理中心
	ctx, cancel := context.WithCancel(context.Background())
	xlog.Info(ctx) // 后续使用ctx则注释, 主协程控制入口
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGABRT, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		cancel()
		// stop server
		model.Close()
	}()

	// router 初始化


	svr.ServerStart()
	xlog.Debug("bye~")
}
