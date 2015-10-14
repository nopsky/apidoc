package main

import (
	"fmt"
	"strings"
)


var docReturnTypeFormat = `

> 返回类型

%s

`

type returnType struct {
	name 	string
} 


type returnTypeToken struct {
	_returnType *returnType
	doc 		string

}

func newReturnTypeToken() (t *returnTypeToken) {
	t = &returnTypeToken{}
	t._returnType = &returnType{}
	return
}

func (self *returnTypeToken) parse(strArr []string) {
	s := make([]string, 2)
	copy(s, strArr)
	//s[0] = @return_type
	self._returnType.name = s[1]
}

//生成macdown文档
func (self *returnTypeToken) makeMacDown() (md string){
	self.doc = fmt.Sprintf(docReturnTypeFormat, strings.ToUpper(self._returnType.name))
	return
}