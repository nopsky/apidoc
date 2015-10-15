package main


import (
	"fmt"
	//"errors"
)


var docReturnParamFormat = `

> 返回字段说明

字段名|变量名|类型|描述
---|---|---|---
%s

`


type returnParam struct {
	name 		string
	varName 	string
	varType		string
	desc		string
}

type returnParamToken struct {
	_returnParam []*returnParam
	doc 		string
}

func newReturnParamToken() (t *returnParamToken) {
	t = &returnParamToken{}
	t._returnParam = make([]*returnParam, 0, 20)
	return
}

func (self *returnParamToken) parse(strArr []string) {
	s := make([]string, 5)
	copy(s, strArr)
	//s[0] = "@return_param"
	t := &returnParam{}
	t.name = s[1]
	t.varName = s[2]
	if s[3] == "" {
		s[3] = "any"
	}

	t.varType = s[3]	
	if s[4] == "" {
		s[4] = "无"
	}
	t.desc = s[4]

	self._returnParam = append(self._returnParam, t);
}

//生成macdown文档
func (self *returnParamToken) makeMacDown(){
	var str string
	if len(self._returnParam) > 0 {
		for _, v := range self._returnParam {
			str += fmt.Sprintf("%s|%s|%s|%s\n", v.name, v.varName, v.varType, v.desc)
		}
		self.doc = fmt.Sprintf(docReturnParamFormat, str)
	}
	return
}