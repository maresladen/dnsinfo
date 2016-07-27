package main

import (
	"archive/tar"
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

//获得IP地址并写文件
func getIPA(htmlAdd string) {
	ns, err := net.LookupHost(htmlAdd)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s", err.Error())
		return
	}
	//TODO 写一个判断文件夹是否存在的方法
	//TODO 这里改为将ip写入到一个文件中
	for _, n := range ns {
		fmt.Fprintf(os.Stdout, "--%s\n", n)
	}
}

func readLine(filename string, handler func(string)) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		handler(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}

//
func zipFolder(sourceFolder, targetFile string) {
	// file write
	fw, err := os.Create(targetFile)
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
		fmt.Println(fi.Name())
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

func unzipFolder(unzipFile string) {
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
		fw, err := os.OpenFile("file2/"+h.Name, os.O_CREATE|os.O_WRONLY, 0644 /*os.FileMode(h.Mode)*/)
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

func untarFile(file, path string) error {
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
