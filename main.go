package main

import (
	"fmt"
	"os"
	"strings"
	tf "toolFolder"
)

func main() {

	argNum := len(os.Args)
	if argNum > 2 || argNum <= 1 {
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

	tf.ConfigSet()
	// fmt.Println(tf.GetEnvPath() + tf.DnsFileName)
	//建立文件夹
	// tf.CreateFloder(tf.Constfolder)
	//委托方法，通过此方法建立文件
	tempFun := tf.GetIPA
	//读取配置文件，并调用委托方法
	tf.ReadLine("list.txt", tempFun)
	//压缩文件,不再压缩文件,直接上传
	// tf.ZipFile(tf.DnsFileName, "dp.tar.gz")
	//上传到7牛，设定一个独立的id号
	tf.SevenDelFile(tf.DnsFileName)
	tf.SevenCreateFile(tf.GetEnvPath()+tf.DnsFileName, tf.DnsFileName)
}

func clientFun() {
	tf.ConfigSet()
	htmlAddr := tf.SevenGetDownLoadUrl()
	fmt.Println(htmlAddr)
	tf.DownloadFiles(htmlAddr)
	fmt.Println("下载完成??")
	// tf.UntarFile("dp.tar.gz", tf.DnsFilePath)
	fmt.Println("dnsmasq重启,自动运行命令行")
}
