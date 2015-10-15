package main


import (
	//"fmt"
	"container/list"
)


type Parse struct {
	class 			*classToken
	name 			*nameToken
	author			*authorToken
	method 			*methodToken 
	input			*inputParamToken
	returnType		*returnTypeToken
	returnParam 	*returnParamToken
	exmaple 		*exmapleToken
}


func newParse() (p *Parse){
	p = &Parse{}
	p.class = newClassToken()
	p.name = newNameToken()
	p.author = newAuthorToken()
	p.method = newMethodToken()
	p.input = newInputParamToken()
	p.returnType = newReturnTypeToken()
	p.returnParam = newReturnParamToken()
	p.exmaple = newExmapleToken()
	return
}

func (p *Parse) parseToken(tokenList list.List) {
	for e := tokenList.Front(); e != nil; e = e.Next() {
		strArr := e.Value.([]string);
		switch strArr[0] {
		case "@class":
			p.class.parse(strArr)
		case "@name":
			p.name.parse(strArr)
		case "@author":
			p.author.parse(strArr)
		case "@method":
			p.method.parse(strArr)
		case "@input_param":
			p.input.parse(strArr)
		case "@return_type":
			p.returnType.parse(strArr)
		case "@return_param":
			p.returnParam.parse(strArr)
		case "@exmaple":
			p.exmaple.parse(strArr)
		default:
			//fmt.Println("未处理类型:", strArr[0])
		}
	}
}

func (p *Parse) makeMacDown(path string)(className string, classPath string, interfaceName string, doc string) {
	p.name.makeMacDown()
	doc += p.name.doc
	
	//生成调用地址的macdown,并获取生成文档的路径
	p.method.makeMacDown()
	doc += p.method.doc
	
	p.input.makeMacDown()
	doc += p.input.doc
	
	p.returnType.makeMacDown()
	doc += p.returnType.doc

	p.exmaple.makeMacDown()
	doc += p.exmaple.doc
	
	p.returnParam.makeMacDown()
	doc += p.returnParam.doc
	
	p.author.makeMacDown()
	doc += p.author.doc

	return p.class.className, p.class.classPath, p.name._apiInfo.name,  doc

}


