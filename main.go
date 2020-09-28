package main

import (
	"archive/zip"
	"fmt"
	"github.com/atotto/clipboard"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	client := &http.Client{}
	res,err:= client.Get("http://idea.medeming.com/jets/images/jihuoma.zip")

	if err != nil {
		panic("下载失败")
	}
	bytes,err := ioutil.ReadAll(res.Body)


	err = ioutil.WriteFile("./jihuoma.zip",bytes,os.ModePerm)
	if err != nil {
		panic(err.Error())
	}

	r,e :=zip.OpenReader("./jihuoma.zip")
	if e != nil {
		panic("解压失败")
	}
	defer r.Close()
	for _,f:= range r.File{
		rc,err:= f.Open()
		if err != nil {
			panic("解压失败")
		}
		defer rc.Close()
		//fmt.Println(f.Name)
		if !strings.Contains(f.Name,"later.txt"){
			continue
		}
		bytes,err:= ioutil.ReadAll(rc)
		if err != nil {
			panic("解压失败")
		}
		clipboard.WriteAll(string(bytes))
		fmt.Println(string(bytes))
		os.Remove("./jihuoma.zip")
	}
}

