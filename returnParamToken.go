package main


import (
	"fmt"
	"github.com/Jeffail/gabs"
	"strings"
	//"errors"
)

var docReturnTextFormat = "> 返回结果\n\n```\n%s\n```\n"

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
	parent 		string     //父类
	desc		string
	varValue	interface{}

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
	s := make([]string, 6)
	copy(s, strArr)
	//s[0] = "@return_param"
	t := &returnParam{}
	t.name = s[1]
	t.varName = s[2]
	if s[3] == "" {
		s[3] = "any"
	} else {
		switch(strings.ToLower(s[3])) {
		case "int":
			t.varValue = 100
		case "string":
			t.varValue = "nopsky"
		case "float" :
			t.varValue = 100.00
		case "url" :
			t.varValue = "http://github.com/nopsky"
		case "bool" :
			t.varValue = true
		case "array+int":
			t.varValue = []int{1,2,3,4,5}
		case "array+string":
			t.varValue = []string{"a", "b", "c", "d"}
		case "array+float":
			t.varValue = []float32{1.0, 2.0, 3.0, 4.0, 5.0}
		case "array+url":
			t.varValue = []string{"http://www.google.com", "http://www.stackoverflow.com", "http://gihub.com/nopsky"}
		default:
			//errors.New("未知的类型:" + s[3])
			fmt.Println("未知的类型:", s[3])
			return
		}
	}
	t.varType = s[3]	
	t.parent = strings.ToUpper(s[4])
	if s[5] == "" {
		s[5] = "无"
	}
	t.desc = s[5]

	self._returnParam = append(self._returnParam, t);
}

//生成macdown文档
func (self *returnParamToken) makeMacDown(){
	var str string

	jsonObj := gabs.New()
	for _, v := range self._returnParam {
		str += fmt.Sprintf("%s|%s|%s|%s\n", v.name, v.varName, v.varType, v.desc)

		var path string 
		if v.parent == "" || v.parent == "NULL" {
			path = v.varName
		} else {
			path = v.parent + "." + v.varName
		}
		switch(v.varType) {
			case "int":
				fallthrough
			case "string":
				fallthrough
			case "float" :
				fallthrough
			case "url" :
				fallthrough
			case "bool" :
				jsonObj.SetP(v.varValue, path)
			case "array+int":
				jsonObj.ArrayP(path)
				for _, vv := range v.varValue.([]int) {
					jsonObj.ArrayAppendP(vv, path)
				}
			case "array+string":
				fallthrough
			case "array+url":
				jsonObj.ArrayP(path)
				for _, vv := range v.varValue.([]string) {
					jsonObj.ArrayAppendP(vv, path)
				}
			case "array+float":
				jsonObj.ArrayP(path)
				for _, vv := range v.varValue.([]float32) {
					jsonObj.ArrayAppendP(vv, path)
				}		
			}
	}
	self.doc = fmt.Sprintf(docReturnTextFormat, jsonObj.StringIndent("", "  ")) + fmt.Sprintf(docReturnParamFormat, str)
	return
}