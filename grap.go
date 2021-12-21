package lad

import (
	"container/list"
	"fmt"
)

type node struct {
	val string
}

func newGraphNode(val string) *node {
	return &node{
		val: val,
	}
}

type tokenGraph struct {
	nodes map[*node]*list.List
	start *node
}

func newTokenGraph() *tokenGraph {
	start := &node{
		val: "/",
	}

	rs := &tokenGraph{
		nodes: make(map[*node]*list.List),
		start: start,
	}

	return rs
}

func (tg *tokenGraph) bfs() {
	l := make([]*node, 0)
	fmt.Println(tg.start.val)
	visited := make(map[*node]struct{})
	l = append(l, tg.start)
	visited[tg.start] = struct{}{}
	for len(l) > 0 {
		n := l[0]
		l = l[1:]
		_, ok := tg.nodes[n]
		if ok && tg.nodes[n].Len() > 0 {
			tn := tg.nodes[n].Front()
			tnv := tn.Value.(*node)
			if _, ok := visited[tnv]; !ok {
				fmt.Println(tnv.val)
				visited[tnv] = struct{}{}
				l = append(l, tnv)
			}

			for tn := tn.Next(); tn != nil; tn = tn.Next() {

				tnv := tn.Value.(*node)
				if _, ok := visited[tnv]; !ok {
					fmt.Println(tnv.val)
					visited[tnv] = struct{}{}
					l = append(l, tnv)
				}
			}
		}
	}
}

func (tg *tokenGraph) dfs() {
	l := make([]*node, 0)
	visited := make(map[*node]struct{})
	l = append(l, tg.start)
	for len(l) > 0 {
		n := l[len(l)-1]
		l = l[0 : len(l)-1]

		_, ok := tg.nodes[n]
		if !ok {
			continue
		}

		tn := tg.nodes[n].Front()
		for tn != nil {
			tnv := tn.Value.(*node)
			if _, ok := visited[tnv]; !ok {
				visited[tnv] = struct{}{}
				fmt.Println(tnv.val)
				l = append(l, n, tnv)
				break
			}
			tn = tn.Next()
		}
	}
}

func (tg *tokenGraph) appendEdge(s *node, e ...*node) {
	if _, ok := tg.nodes[s]; !ok {
		tg.nodes[s] = list.New()
	}
	for _, end := range e {
		tg.nodes[s].PushBack(end)
	}
}
