package main

import (
	"fmt"
)

var docAuthorFormat = `

> 作者

%s  <%s>

`

type author struct {
	username 	string 
	email 		string
}

type authorToken struct {
	_author		*author
	doc 		string
}

func newAuthorToken() (t *authorToken) {
	t = &authorToken{}
	t._author = &author{}
	return
}

func (self *authorToken) parse(strArr []string) {
	s := make([]string, 3)
	copy(s, strArr)
	//s[0] = @author
	self._author.username = s[1]
	self._author.email = s[2]
}

//生成macdown文档
func (self *authorToken) makeMacDown(){
	self.doc = fmt.Sprintf(docAuthorFormat, self._author.username, self._author.email)
	return
}
