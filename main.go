package main


import (
	"io/ioutil"
	"os"
	"sync"
	"path/filepath"
	"strings"
	"fmt"
	"flag"
)

//文件列表
var fileList []string 

var wg sync.WaitGroup

//文件后缀
var suffix = flag.String("suffix", "", "文件后缀, 例如:.go")

var path = flag.String("path", "", "搜索目录, 例如:/data/work/")

var docDir = flag.String("doc", ".", "生成文档的目录, 例如:/data/doc/")

func listFunc(path string, f os.FileInfo, err error) error {

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
			go func(file string) {
				defer wg.Done()
				file_contents := readFile(file)

				lexer := newLexer()
				lexer.startLexer(file_contents)

				p := newParse()
				p.parseToken(lexer.tokenList)

				className, classPath, name, apiDocText := p.makeMacDown(*docDir)

				if name == "" {
					return;
				}
				//api文档文件的地址
				docFile := docPath + "/" + classPath + "/" + name + ".md"

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


				//创建api文件
				err = ioutil.WriteFile(docFile, []byte(apiDocText), 0777)

				if err != nil {
					panic(err)
				}

				if className != "" {
					if _, ok := index[className]; !ok {
						index[className] = make(map[string]string)
					}

					index[className][name] = classPath + "/" + name + ".md"
				}
		    }(file)
		}
	} else {
		fmt.Println("没有符合条件的文件")
	}

	wg.Wait()
	//根据生成的API接口文档，生成索引文件
	
	if len(index) > 0 {
		var indexContent = "##API 接口文档\n\n"
		for k, v := range index {
			indexContent += "###"+k+"\n\n"
			for kk, vv := range v {
				indexContent += "* ["+kk+"]("+vv+")\n\n"
			}
		}

		//创建索引文件
		err = ioutil.WriteFile(docPath + "/" + "index.md", []byte(indexContent), 0777)

		if err != nil {
			panic(err)
		}
	} else  {
		fmt.Println("没有符合格式的文档")
	}

}