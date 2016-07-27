package toolFolder

import (
	"fmt"

	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodo"
	"qiniupkg.com/api.v7/kodocli"
)

var (
	//设置上传到的空间
	bucket = "cooldan"
	//设置上传文件的key
	// key = "dp.tar.gz"
	key = "FvZBWWp3oAFqfHYSrFTnPIDSnyzU"

	//acckey
	acckey = "rnoW5udV8fWKnwqFIOKxclr-V52DTajPLmvDJnlD"
	//aeckey
	seckey = "hjOYvRIgClcIS4grZ00UmkjflxknyJvFnENwCTFo"
)

//PutRet 构造返回值字段
type PutRet struct {
	Hash string `json:"hash"`
	Key  string `json:"key"`
}

//CreateFile 建立文件，会返回七牛自动哈希的ID
func CreateFile(filepath string) {
	//初始化AK，SK
	conf.ACCESS_KEY = acckey
	conf.SECRET_KEY = seckey
	//创建一个Client
	c := kodo.New(0, nil)
	//设置上传的策略
	policy := &kodo.PutPolicy{
		Scope: bucket,
		//设置Token过期时间
		Expires: 3600,
	}
	//生成一个上传token
	token := c.MakeUptoken(policy)

	//构建一个uploader
	zone := 0
	uploader := kodocli.NewUploader(zone, nil)

	var ret PutRet
	//设置上传文件的路径
	// filepath := "/Users/dxy/sync/sample2.flv"
	//调用PutFileWithoutKey方式上传，没有设置saveasKey以文件的hash命名
	res := uploader.PutFileWithoutKey(nil, &ret, token, filepath, nil)
	//打印返回的信息
	fmt.Println(ret)
	//打印出错信息
	if res != nil {
		fmt.Println("io.Put failed:", res)
		return
	}
}

//ConverFile 覆盖文件，需要提供要覆盖的文件的ID，在7牛上传这个文件的时候，本身就定义好ID，这样的话以后就不会有问题
func ConverFile(filepath string) {
	//初始化AK，SK
	conf.ACCESS_KEY = acckey
	conf.SECRET_KEY = seckey
	//创建一个Client
	c := kodo.New(0, nil)

	//设置上传的策略
	policy := &kodo.PutPolicy{
		Scope: bucket + ":" + key,
		//设置Token过期时间
		Expires: 3600,
	}
	//生成一个上传token
	token := c.MakeUptoken(policy)

	//构建一个uploader
	zone := 0
	uploader := kodocli.NewUploader(zone, nil)

	var ret PutRet
	//设置上传文件的路径
	// filepath := "/Users/dxy/sync/sample2.flv"
	//调用PutFile方式上传，这里的key需要和上传指定的key一致
	res := uploader.PutFile(nil, &ret, token, key, filepath, nil)
	//打印返回的信息
	fmt.Println(ret)
	//打印出错信息
	if res != nil {
		fmt.Println("io.Put failed:", res)
		return
	}
}
