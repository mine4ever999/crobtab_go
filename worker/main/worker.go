package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"

	"go-crontab/worker"
)

var (
	confFile string // 配置文件路径
)

// 解析命令行参数
func initArgs() {
	// worker -config ./worker.json
	// worker -h
	flag.StringVar(&confFile, "config", "./worker.json", "worker.json")
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

	// 初始化命令行参数
	initArgs()

	// 初始化线程
	initEnv()

	// 初始化配置文件
	if err = worker.InitConfig(confFile); err != nil {
		goto ERR
	}

	// 初始化任务管理器
	if err = worker.InitJobMgr(); err != nil {
		goto ERR
	}

	// 初始化任务调度器
	if err = worker.InitScheduler(); err != nil {
		goto ERR
	}

	// 初始化任务执行器
	if err = worker.InitExecutor(); err != nil {
		goto ERR
	}

	// 正常退出
	for {
		time.Sleep(1 * time.Second)
	}
ERR:
	fmt.Println(err)
}
