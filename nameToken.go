package main


import (
	"fmt"
)

var docNameFormat = `

> 接口名称

%s

> 调用地址

%s

`

type apiInfo struct {
	name 		string
	route 		string
}

type nameToken struct {
	_apiInfo	*apiInfo
	doc 		string
}

func newNameToken() (t *nameToken) {
	t = &nameToken{}
	t._apiInfo = &apiInfo{}
	return
}

func (self *nameToken) parse(strArr []string) {
	s := make([]string, 3)
	copy(s, strArr)
	//s[0] = @name
	self._apiInfo.name = s[1]
	self._apiInfo.route = s[2]
}

//生成macdown文档
func (self *nameToken) makeMacDown(){
	self.doc = fmt.Sprintf(docNameFormat, self._apiInfo.name, self._apiInfo.route)
	return
}

