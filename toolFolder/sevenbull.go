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
	//acckey
	acckey = "rnoW5udV8fWKnwqFIOKxclr-V52DTajPLmvDJnlD"
	//aeckey
	seckey = "hjOYvRIgClcIS4grZ00UmkjflxknyJvFnENwCTFo"

	domain = "7xoyml.com1.z0.glb.clouddn.com"
)

//PutRet 构造返回值字段
type PutRet struct {
	Hash string `json:"hash"`
	Key  string `json:"key"`
}

//SevenCreateFile 建立文件，会返回七牛自动哈希的ID
func SevenCreateFile(filepath, filename string) {
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
	err := uploader.PutFile(nil, &ret, token, filename, filepath, nil)
	//打印出错信息
	if err != nil {
		writelog(err, "七牛建立文件错误")
		return
	}

}

//SevenConverFile 覆盖文件，需要提供要覆盖的文件的ID，在7牛上传这个文件的时候，本身就定义好ID，这样的话以后就不会有问题
func SevenConverFile(filepath string) {
	//初始化AK，SK
	conf.ACCESS_KEY = acckey
	conf.SECRET_KEY = seckey
	//创建一个Client
	c := kodo.New(0, nil)

	//设置上传的策略
	policy := &kodo.PutPolicy{
		Scope: bucket + ":" + DnsFileName,
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
	res := uploader.PutFile(nil, &ret, token, DnsFileName, filepath, nil)
	//打印返回的信息
	fmt.Println(ret)
	//打印出错信息
	if res != nil {
		writelog(res, "上传失败")
		return
	}

}

func SevenGetDownLoadUrl() string {
	//初始化AK，SK
	conf.ACCESS_KEY = acckey
	conf.SECRET_KEY = seckey
	//调用MakeBaseUrl()方法将domain,key处理成http://domain/key的形式
	baseUrl := kodo.MakeBaseUrl(domain, DnsFileName)
	policy := kodo.GetPolicy{}
	//生成一个client对象
	c := kodo.New(0, nil)
	//调用MakePrivateUrl方法返回url
	return c.MakePrivateUrl(baseUrl, &policy)
}

func SevenDelFile(strKey string) {
	c := kodo.New(0, nil)
	p := c.Bucket(bucket)

	//调用Delete方法删除文件
	res := p.Delete(nil, strKey)
	//打印返回值以及出错信息
	if res == nil {
		fmt.Println("Delete success")
	} else {
		fmt.Println(res)
	}
}
