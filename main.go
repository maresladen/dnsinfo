package main

import (
	"fmt"
	"os"
	"strings"
	tf "toolFolder"
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

	// fmt.Println("hello dns")
	//---------------------------暂时屏蔽----------------------
	argNum := len(os.Args) 
	if argNum > 2  || argNum <= 1{
		fmt.Println("只定义一个参数，分别为[-S]服务端,[-C]客户端")
	}

	if argNum == 2 {
		switch strings.ToUpper(os.Args[1]) {
		case "-S":
			fmt.Println("服务端运行")
			serverFun()
			fmt.Println("运行完成")
		case "-C":
			fmt.Println("客户端运行")
			clientFun()
		}

	}
	//---------------------------暂时屏蔽----------------------
	// argNum := len(os.Args)
	// if argNum == 1 {
	// 	fmt.Println("服务端运行")
	// 	serverFun()
	// 	fmt.Println("运行完成")
	// }

}

func serverFun() {
	
	//建立文件夹
	// tf.CreateFloder(tf.Constfolder)
	//委托方法，通过此方法建立文件
	tempFun := tf.GetIPA
	//读取配置文件，并调用委托方法
	tf.ReadLine("list", tempFun)
	//压缩文件
	tf.ZipFolder("dp","dp.tar.gz")
	//上传到7牛，设定一个独立的id号
	tf.SevenConverFile("dp.tar.gz")
}

func clientFun() {
	fmt.Println("我是客户端，实际方法还未写")
}
