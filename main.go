package main

import (
	"fmt"
)

func main() {

	//1 抓取gfw的屏蔽内容,txt 这里手工做，不需要每次都去抓，偶尔维护一下就可以了
	//2 通过屏蔽内容的txt 得到Aip地址 
	//3 生成文件
	//4 ok 打包
	//5 ok 推到7牛
	//6 定时服务 从7牛下载 这个可能要写shell，运行，运行完之后关闭
	//7 保存 在上面一同做掉
	//8 重启bind 这个在程序运行完之前，在程序中执行shell命令

	fmt.Println("hello dns")
}
