package main

import (
	"fmt"
	"strings"
)


var docMethodFormat = `

> 调用方式

%s

`  
type method struct {
	name 	string
}

type methodToken struct {
	_method		*method
	doc 		string
}

func newMethodToken() (t *methodToken) {
	t = &methodToken{}
	t._method = &method{}
	return
}

func (self *methodToken) parse(strArr []string) {
	s := make([]string, 2)
	copy(s, strArr)
	//s[0] = @method
	self._method.name = s[1]
}

//生成macdown文档
func (self *methodToken) makeMacDown() {
	if self._method.name != "" {
		self.doc = fmt.Sprintf(docMethodFormat, strings.ToUpper(self._method.name))
	} else {
		self.doc = ""
	}
}