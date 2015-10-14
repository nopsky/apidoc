package main


import (
	"fmt"
	"strings"
)

var docInputParamFormat = `

> 参数


字段名|变量名|类型|必填|示例值|描述
---|---|---|---|---|---
%s

`

type inputParam struct {
	name 		string 		//字段名
	varName		string 		//变量名
	varType		string 		//变量类型
	isMust		string		//是否必须
	ex 			string      //示例值
	desc		string  	//描述
}

type inputParamToken struct {
	_inputParam 	[]*inputParam
	doc 			string      //生成的文档
}

func newInputParamToken() (t *inputParamToken) {
	t = &inputParamToken{}
	t._inputParam = make([]*inputParam, 0, 10)
	return
}

func (self *inputParamToken) parse(strArr []string) {
	s := make([]string, 7)
	copy(s, strArr)
	input := &inputParam{}
	input.name = s[1]
	input.varName = s[2]
	input.varType = strings.ToLower(s[3])
	if s[4] == "" || s[4] == "0" {
		s[4] = "false"
	} else {
		s[4] = "true"
	}
	input.isMust = s[4]
	//如果没有示例，根据变量类型，默认给出一个示例值
	if s[5] == "" || strings.ToUpper(s[5]) == "NULL"{
		switch(input.varType) {
		case "string" :
			s[5] = "nopsky"
		case "int" :
			s[5] = "100"
		case "bool" :
			s[5] = "true"
		case "url" :
			s[5] = "http://github.com/nopsky"
		case "float" :
			s[5] = "100.00"
		case "array" :
			s[5] = "[1,2,3,4,5]"
		case "file" :
			s[5] = "fileName"
		}
	}
	input.ex = s[5]
	if s[6] == "" {
		s[6] = "无"
	}
	input.desc = s[6]
	self._inputParam = append(self._inputParam, input)
}

//生成macdown文档
func (self *inputParamToken) makeMacDown() {
	var str string
	if len(self._inputParam) > 0 {
		for _, v := range self._inputParam {
			//fmt.Println(v)
			str += fmt.Sprintf("%s|%s|%s|%s|%s|%s\n", v.name, v.varName, v.varType, v.isMust, v.ex, v.desc)
		}
		self.doc = fmt.Sprintf(docInputParamFormat, str)		
	} else {
		self.doc = ">参数\n\n无\n"
	}

	return
}