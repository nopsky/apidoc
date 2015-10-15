package main

import (
	"fmt"
	"github.com/jeffail/gabs"
)

var exmapleFormat = "\n > 返回结果\n\n```\n%s\n```\n"

type exmaple struct {
	jsonStr		string
}

type exmapleToken struct {
	_exmaple	*exmaple
	doc 		string
}

func newExmapleToken() (t *exmapleToken) {
	t = &exmapleToken{}
	t._exmaple = &exmaple{}
	return
}

func (self *exmapleToken) parse(strArr []string) {
	s := make([]string, 2)
	copy(s, strArr)
	//s[0] = @author
	self._exmaple.jsonStr = s[1]
}

//生成macdown文档
func (self *exmapleToken) makeMacDown(){
	
	if self._exmaple.jsonStr != "" {
		jsonParsed, err := gabs.ParseJSON([]byte(self._exmaple.jsonStr))
		
		if err != nil {
			fmt.Println("示例的JSON格式不正确")
			return
		}

		self.doc = fmt.Sprintf(exmapleFormat, jsonParsed.StringIndent("", " "))
	}
	return
}