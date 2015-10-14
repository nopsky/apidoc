package main

import (
	"container/list"
	"strings"
)

type Lexer struct {
	currentPos				int 			//当前的偏移量
	tokenList 				list.List 		//存放注释的列表
}

func newLexer() (lexer *Lexer) {
	self := &Lexer{}
	self.currentPos = 0;
	return self
}

func (self *Lexer) startLexer(file_contents string) {
	for self.currentPos < len(file_contents) {
		content := file_contents[self.currentPos:]
		mCommentStartPos := strings.Index(content, "/*")
		sCommentStartPos := strings.Index(content, "//")

		if(sCommentStartPos == -1 && mCommentStartPos == -1) {
			break
		}

		//如果找到了注释符
		if mCommentStartPos > sCommentStartPos {
			if sCommentStartPos >= 0 {
				//处理单行注释
				content = content[sCommentStartPos:]
				endPos := self.parse(content)
				self.currentPos = self.currentPos+endPos+1
			} else {
				//处理多行注释
				mCommentEndPos := strings.Index(content, "*/")
				content = content[mCommentStartPos:mCommentEndPos]
				for{
					endPos := self.parse(content)
					if endPos == len(content) {
						break;
					}
					content = content[endPos+1:]
				}
				self.currentPos = self.currentPos+ mCommentEndPos + 2

			}
		} else {
			if mCommentStartPos >= 0 {
				//处理多行注释
				mCommentEndPos := strings.Index(content, "*/")
				content = content[mCommentStartPos:mCommentEndPos]
				for{
					endPos := self.parse(content)
					if endPos == len(content) {
						break;
					}
					content = content[endPos+1:]
				}
				self.currentPos = self.currentPos + mCommentEndPos + 2

			} else {
				//处理单行注释
				content = content[sCommentStartPos:]
				endPos := self.parse(content)
				self.currentPos = self.currentPos + endPos + 1
			}
		}
	}
}

func (self *Lexer) getToken(content string) []string{
	start := strings.Index(content, "@")
	if start == -1 {
		return nil
	}

	return strings.Fields(content[start:]);
}

func (self *Lexer) parse(content string) (pos int) {
	pos = strings.Index(content, "\n")
	if pos == -1 {
		pos = len(content)
		
	}
	sContent := content[:pos]
	list := self.getToken(sContent)
	if list != nil {
		self.tokenList.PushBack(list)
	}
	return pos
}