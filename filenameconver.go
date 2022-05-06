package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)
var (
	isContinue =true
	path string
	oldSuffix string
	newSuffix string
	subFolderTrans bool
)
func main()  {
	for isContinue {
		var selectNum int
		fmt.Printf(" \n")
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>菜单Start<<<<<<<<<<<<<<<<<<<<<<<<")
		fmt.Println("1、请输入需要转换后缀文件的文件夹绝对路径： 		 	此时为：",path)
		fmt.Println("2、请输入需要转换的后缀： 						 	此时为：",oldSuffix)
		fmt.Println("3、请输入需要转换成的后缀：						 	此时为：",newSuffix)
		fmt.Println("4、是否子文件夹中也全部转换,请答true或false 	     	此时为：",subFolderTrans)
		fmt.Println("5、开始转换")
		fmt.Println("6、关闭程序")
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>菜单End<<<<<<<<<<<<<<<<<<<<<<<<<<")
		fmt.Printf(" \n")
		fmt.Scanln(&selectNum)
		switch selectNum {
		case 1:
			fmt.Println("请输入需要转换后缀文件的文件夹绝对路径：")
			fmt.Scanln(&path)
		case 2:
			fmt.Println("请输入需要转换的后缀：")
			fmt.Scanln(&oldSuffix)
		case 3:
			fmt.Println("请输入需要转换成的后缀：")
			fmt.Scanln(&newSuffix)
		case 4:
			fmt.Println("是否子文件夹中也全部转换,请答true或false ")
			_,err:= fmt.Scanln(&subFolderTrans);if err!=nil{
			fmt.Println(err)
		}

		case 5:

			if path=="" {
				fmt.Println("路径为空")
				continue
			}
			if oldSuffix==""{
				fmt.Println("需要转换的后缀为空")
				continue
			}

			if newSuffix==""{
				fmt.Println("需要转换成的后缀为空")
				continue
			}
			translate(path)
		case 6:
			isContinue=false
		}
	}

}

func translate(path string)  {

	transNum:=0
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	}
	if len(files)==0{
		fmt.Println(path,"路径文件为 0,")
	}
	for _, f := range files {
		fileName:=f.Name()
		if f.IsDir() && subFolderTrans{
			translate(path +"\\" + fileName)
		}
		suffix:=strings.Split(fileName,".")
		if suffix[len(suffix)-1]==oldSuffix{
			transNum++
			fmt.Println("开始转换文件:",path +"\\" + fileName)
			suffix[len(suffix)-1]=newSuffix
			os.Rename(path +"\\" + fileName, path +"\\" + strings.Join(suffix,"."))
			fmt.Println("成功转换为换为:",path +"\\" + strings.Join(suffix,"."))
		}
	}
	fmt.Println(path,"路径下", transNum, "个文件  全部转换完成！！")
}