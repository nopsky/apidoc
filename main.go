package main


import (
	"io/ioutil"
	"os"
	//"io"
	"sync"
	"path/filepath"
	"strings"
	"fmt"
	"flag"
)

//获取系统类型
//var osType = os.Getenv("GOOS")

//文件列表
var fileList []string 

var wg sync.WaitGroup

//文件后缀
var suffix = flag.String("suffix", "", "文件后缀, 例如:.go")

var path = flag.String("path", "", "搜索目录, 例如:/data/work/")

var docDir = flag.String("doc", ".", "生成文档的目录, 例如:/data/doc/")

func listFunc(path string, f os.FileInfo, err error) error {
	// var dirSep string

	// if osType == "windows" {
	// 	dirSep = "\\"
	// } else if osType == "linux" {
	// 	dirSep = "/"
	// }

	if f == nil {
		return err
	}

	if f.IsDir() {
		return nil
	}

	if *suffix != "" {
		ok := strings.HasSuffix(path, *suffix)

		if !ok {
			return nil
		}
	} 

	fileList = append(fileList, path)
	return nil
}

//获取文件列表
func getFileList(path string) {
	err := filepath.Walk(path, listFunc)

	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

func readFile(fileName string) string {
	b, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	return string(b)
}

func main() {

	flag.Parse()

	if *path == "" {
		flag.PrintDefaults()
		return
	}

	realPath, err := filepath.Abs(filepath.Dir(*path));

	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	docPath, err := filepath.Abs(*docDir);
	
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	//获取API接口文件列表
	getFileList(realPath)
	index := make(map[string]map[string]string)
	//生成API接口文档
	if len(fileList) > 0 {
		for _, file := range fileList {
			wg.Add(1)
			go func() {
				defer wg.Done()

				file_contents := readFile(file)

				lexer := newLexer()
				lexer.startLexer(file_contents)

				p := newParse()
				p.parseToken(lexer.tokenList)
				name, route, apiDocText := p.makeMacDown(*docDir)

				//api文档文件的地址
				docFile := docPath + "/" + route + ".md"

				//api文档目录
				apiDocPath, err := filepath.Abs(filepath.Dir(docFile))

				if err != nil {
					panic(err)
				}

				//创建api文档目录
				err = os.MkdirAll(apiDocPath, 0755)

				if err != nil {
					fmt.Printf("生成目录失败%v\n", err)
					return
				}

				//baseFileName := _path.Base(docFile)

				//docFile = apiDocPath + "/" + name + ".md"

				//创建api文件
				err = ioutil.WriteFile(docFile, []byte(apiDocText), 0777)

				if err != nil {
					panic(err)
				}

				if _, ok := index[route]; !ok {
					index[route] = make(map[string]string)
				}

				index[route][name] = route + ".md"
			}()
		}
	} else {
		fmt.Println("没有符合条件的文件")
	}

	wg.Wait()
	//根据生成的API接口文档，生成索引文件
	//fmt.Println(*index)
	
	var indexContent = "##API 接口文档\n"
	for k, v := range index {
		k = filepath.Dir(k)
		indexContent += "###"+k+"\n"
		for kk, vv := range v {
			indexContent += "* ["+kk+"]("+vv+")"
		}
	}

	//创建索引文件
	err = ioutil.WriteFile(docPath + "/" + "index.md", []byte(indexContent), 0777)

	if err != nil {
		panic(err)
	}

}