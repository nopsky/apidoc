package main

/**
 * 类名和类路径，主要用于生成的接口文档的目录生成和接口索引文件分组
 */
type classToken struct {
	className		string
	classPath 		string
}

func newClassToken() (t *classToken) {
	t = &classToken{}
	return
}

func (self *classToken) parse(strArr []string) {
	s := make([]string, 3)
	copy(s, strArr)
	//s[0] = @author
	self.className = s[1]
	self.classPath = s[2]
}
