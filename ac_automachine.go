package lad

import (
	"bufio"
	"container/list"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

const rootRaw = "/"

type acNode struct {
	raw      string
	children map[string]*acNode
	isEnd    bool
	length   int
	deep     int
	fail     *acNode
}

func newAcNode(raw string, deep int) *acNode {
	return &acNode{
		raw:      raw,
		children: make(map[string]*acNode),
		deep:     deep,
	}
}

func (an *acNode) view() {
	fmt.Println(an.raw, ",", an.deep, ",", an.length)
	if len(an.children) == 0 {
		return
	}
	for _, node := range an.children {
		node.view()
	}
}

type acMachine struct {
	root *acNode
}

func New() *acMachine {
	root := newAcNode(rootRaw, 0)
	return &acMachine{
		root: root,
	}
}

// Add 添加模式串
func (ac *acMachine) Add(pattern string) {
	p := ac.root
	tok := newToken(pattern)
	length := 0
	deep := 1
	for str := tok.next(); str != ""; str = tok.next() {
		if _, ok := p.children[str]; !ok {
			newNode := newAcNode(str, deep)
			p.children[str] = newNode
		}
		p = p.children[str]
		length++
		deep++
	}

	p.length = length
	p.isEnd = true
}

// Load 加载文件
func (ac *acMachine) Load(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}
		ac.Add(line)
	}
}

// Build 构建自动机
func (ac *acMachine) Build() {
	l := list.New()

	l.PushBack(ac.root)

	for l.Len() > 0 {
		e := l.Front()
		l.Remove(e)
		p := e.Value.(*acNode)

		for _, pc := range p.children {
			if p == ac.root {
				pc.fail = ac.root
			} else {
				q := p.fail

				for q != nil {
					if qc, ok := q.children[pc.raw]; ok {
						pc.fail = qc
						break
					}
					q = q.fail
				}

				if q == nil {
					pc.fail = ac.root
				}
			}
			l.PushBack(pc)
		}
	}
}

// Find 查找
func (ac *acMachine) Find(text string) []string {
	rs := make([]string, 0)
	ac.match(text, func(tok *token, node *acNode) {
		rs = append(rs, tok.prevNStr(tok.index, node.length))
	})

	return rs
}

// Match 匹配
func (ac *acMachine) Match(text string) bool {
	rs := false
	ac.match(text, func(tok *token, node *acNode) {
		rs = true
	})
	return rs
}

// Replace 替换
func (ac *acMachine) Replace(text, target string) string {
	rs := ""
	ac.match(text, func(tok *token, node *acNode) {
		if rs == "" {
			rs = string(tok.origin)
		}
		rs = strings.Replace(rs, tok.prevNStr(tok.index, node.length), target, -1)
	})
	return rs
}

func (ac *acMachine) match(text string, fn func(tok *token, node *acNode)) {
	p := ac.root
	tok := newToken(text)
	for {
		str := tok.next()
		if str == "" {
			break
		}
		for {
			if _, ok := p.children[str]; !ok && p != ac.root {
				p = p.fail
				continue
			}
			break
		}

		p = p.children[str]

		if p == nil {
			p = ac.root
		}

		tmp := p
		for tmp != ac.root {
			if tmp.isEnd {
				fn(tok, tmp)
			}
			tmp = tmp.fail
		}
	}
}

func (ac *acMachine) matchPinyin(text string, fn func(tok *token, node *acNode)) {
	tok := newToken(text)
	// graph := tok.buildGraph()
	p := ac.root
	for {
		str := tok.next()
		if str == "" {
			break
		}
		for {
			if _, ok := p.children[str]; !ok && p != ac.root {
				p = p.fail
				continue
			}
			break
		}

		p = p.children[str]

		if p == nil {
			p = ac.root
		}

		tmp := p
		for tmp != ac.root {
			if tmp.isEnd {
				fn(tok, tmp)
			}
			tmp = tmp.fail
		}
	}
}
