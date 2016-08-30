package toolFolder

import (
	"archive/tar"
	"bufio"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	Constfolder      = "dp"
	AddressValue     = ""
	DnsFileName      = "test.txt"
	DnsFilePath      = ""
	DnsGetErrAddress = ""
)

type mconfig struct {
	DnsqAddressPATH string
	DnsqAddressName string
}

//读取配置文件
func ConfigSet() {
	var m mconfig
	fi, err := os.Open("DNSClient.json")
	if err != nil {
		writelog(err, "get config json data wrong")
	} else {
		temp, _ := ioutil.ReadAll(fi)
		json.Unmarshal(temp, &m)
		DnsFileName = m.DnsqAddressName
		DnsFilePath = m.DnsqAddressPATH
	}
}

//GetIPA 获取IP地址，并保存为一个文件，放在固定的文件夹中
func GetIPA(htmlAdd string) {
	ns, err := net.LookupHost(htmlAdd)
	if err != nil {
		writelog(err, "获取dns信息失败")
		DnsGetErrAddress += htmlAdd + "\n"
		return
	}
	//TODO 换成多协程运行，需要启用阻塞
	if len(ns) > 0 {
		AddressValue += `address=/` + htmlAdd + `/` + ns[0] + "\n"
		fmt.Println(htmlAdd)
	}
	// ch <- htmlAdd
}

//ReadLine 读取行内容，并执行获取ip地址的方法，保存方法也放在获取ip地址的方法中
func ReadLine(filename string, handler func(string)) error {
	os.Remove(DnsFileName)
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		writelog(err, "打开文件错误")
		return err
	}
	buf := bufio.NewReader(f)
	// tempIndex := 0
	// read2End := false
	// ch := make(chan string, 4)
	for {
		// if read2End {
		// 	break
		// }
		// tempIndex += 1
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)

		handler(line)
		if err != nil {
			if err == io.EOF {
				writeFile(DnsFileName, AddressValue)
				writeFile("dnsErrAddress", DnsGetErrAddress)
				// read2End = true
				return nil
			}
			continue
		}
	}
	// for i := 0; i < tempIndex; i++ {
	// 	tempAddr := <-ch
	// 	fmt.Println(tempAddr)
	// }

	// return nil
}

//ZipFolder 压缩文件夹
func ZipFolder(sourceFolder, targetFile string) {
	// file write
	fw, err := os.Create(targetFile)
	fw.Chmod(0755)
	if err != nil {
		panic(err)
	}
	defer fw.Close()
	// gzip write
	gw := gzip.NewWriter(fw)
	defer gw.Close()
	// tar write
	tw := tar.NewWriter(gw)
	defer tw.Close()
	// 打开文件夹
	dir, err := os.Open(sourceFolder)
	if err != nil {
		panic(nil)
	}
	defer dir.Close()
	// 读取文件列表
	fis, err := dir.Readdir(0)
	if err != nil {
		panic(err)
	}
	// 遍历文件列表
	for _, fi := range fis {
		// 逃过文件夹, 我这里就不递归了
		if fi.IsDir() {
			continue
		}
		// 打印文件名称
		//fmt.Println(fi.Name())
		// 打开文件
		fr, err := os.Open(dir.Name() + "/" + fi.Name())
		if err != nil {
			panic(err)
		}
		defer fr.Close()
		// 信息头
		h := new(tar.Header)
		h.Name = fi.Name()
		h.Size = fi.Size()
		h.Mode = int64(fi.Mode())
		h.ModTime = fi.ModTime()
		// 写信息头
		err = tw.WriteHeader(h)
		if err != nil {
			panic(err)
		}
		// 写文件
		_, err = io.Copy(tw, fr)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("tar.gz ok")
}

func ZipFile(fileName, targetFile string) {
	fw, err := os.Create(targetFile)
	fw.Chmod(0755)
	if err != nil {
		panic(err)
	}
	defer fw.Close()
	// gzip write
	gw := gzip.NewWriter(fw)
	defer gw.Close()
	// tar write
	tw := tar.NewWriter(gw)
	defer tw.Close()
	fr, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer fr.Close()
	fi, err := fr.Stat()
	if err != nil {
		panic(err)
	}
	// 信息头
	h := new(tar.Header)
	h.Name = fi.Name()
	h.Size = fi.Size()
	h.Mode = int64(fi.Mode())
	h.ModTime = fi.ModTime()
	// 写信息头
	err = tw.WriteHeader(h)
	if err != nil {
		panic(err)
	}
	// 写文件
	_, err = io.Copy(tw, fr)
	if err != nil {
		panic(err)
	}
	fmt.Println("tar.gz ok")
}

//UnzipFolder 解压文件夹
func UnzipFolder(unzipFile string) {
	// file read
	fr, err := os.Open(unzipFile)
	if err != nil {
		panic(err)
	}
	defer fr.Close()
	// gzip read
	gr, err := gzip.NewReader(fr)
	if err != nil {
		panic(err)
	}
	defer gr.Close()
	// tar read
	tr := tar.NewReader(gr)
	// 读取文件
	for {
		h, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		// 显示文件
		fmt.Println(h.Name)
		// 打开文件
		fw, err := os.OpenFile(Constfolder+"/"+h.Name, os.O_CREATE|os.O_WRONLY, 0644 /*os.FileMode(h.Mode)*/)
		if err != nil {
			panic(err)
		}
		defer fw.Close()
		// 写文件
		_, err = io.Copy(fw, tr)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("un tar.gz ok")
}

//UntarFile 解压文件
func UntarFile(file, path string) error {
	// 打开文件
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	// 读取GZIP
	gr, err := gzip.NewReader(f)
	if err != nil {
		return err
	}
	defer gr.Close()
	// 读取TAR
	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		if hdr.FileInfo().IsDir() {
			os.MkdirAll(path+string(os.PathSeparator)+hdr.Name, hdr.FileInfo().Mode())
		} else {
			fw, err := os.OpenFile(path+string(os.PathSeparator)+hdr.Name, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, hdr.FileInfo().Mode())
			if err != nil {
				return err
			}
			defer fw.Close()
			_, err = io.Copy(fw, tr)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

//CreateFloder 查看文件夹是否存在，不存在则建立文件夹
func CreateFloder(fName string) {
	err := os.Chdir(fName)
	if err != nil {
		os.Mkdir(fName, 0777)
	}
}

//将内容写入到文件中
func writeFile(filename, strContent string) {

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		writelog(err, "打开文件失败")
	}
	defer file.Close()
	n, err := file.WriteString(strContent)
	if err != nil && n < len(strContent) {
		writelog(err, "写文件失败")
	}
}

//检查文件是否存在
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func writelog(err error, strDefine string) {
	file, _ := os.OpenFile("errlog", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	defer file.Close()
	io.WriteString(file, err.Error()+"  |  "+strDefine+"\n\r")
}

func DownloadFiles(urls string) {
	res, err := http.Get(urls)
	if err != nil {
		writelog(err, "获取七牛地址失败")
		return
	}
	defer res.Body.Close()
	file, err := os.OpenFile(DnsFilePath+DnsFileName, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		writelog(err, "建立文件失败")
		return
	}
	defer file.Close()
	io.Copy(file, res.Body)
}

func GetEnvPath() string {
	file, _ := exec.LookPath(os.Args[0])
	//得到全路径，比如在windows下E:\\golang\\test\\a.exe
	path, _ := filepath.Abs(file)
	//将全路径用\\分割，得到4段，①E: ②golang ③test ④a.exe
	splitstring := strings.Split(path, "\\")

	//size为4
	size := len(splitstring)

	//将全路径用最后一段(④a.exe)进行分割，得到2段，①E:\\golang\\test\\ ②a.exe
	splitstring = strings.Split(path, splitstring[size-1])

	//将①(E:\\golang\\test\\)中的\\替换为/，最终得到结果E:/golang/test/
	rst := strings.Replace(splitstring[0], "\\", "/", size-1)
	return rst
}
