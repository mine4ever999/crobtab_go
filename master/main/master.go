package main

import (
	"flag"
	"fmt"
	"go-crontab/master"
	"runtime"
	"time"
)

var (
	confFile string // 配置文件路径
)

// 解析命令行参数
func initArgs() {
	// master -config ./master.json -xxx 123 -yyy ddd
	// master -h
	flag.StringVar(&confFile, "config", "./master.json", "指定master.json")
	flag.Parse()
}

// 初始化线程数量
func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var (
		err error
	)
	// 0. 初始化命令行参数
	initArgs()

	// 1. 初始化线程
	initEnv()

	// 2. 加载配置
	if err = master.InitConfig(confFile); err != nil {
		goto ERR
	}
	// 3. 初始化任务管理器
	if err = master.InitJobMgr(); err != nil {
		goto ERR
	}
	// 4. 启动API HTTP服务
	if err = master.InitApiServer(); err != nil {
		goto ERR
	}
	// 5. 启动任务调度
	// 6. 启动监听
	// 7. 启动工作节点
	// 8. 启动日志收集
	// 9. 启动监控
	// 10. 启动管理平台
	// 正常退出
	for {
		time.Sleep(1 * time.Second)
	}
ERR:
	fmt.Println(err)
}
