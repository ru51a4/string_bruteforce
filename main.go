package main

import "fmt"

type node struct {
	ch       string
	children []*node
}

func main() {
	_string := "abcdez"
	var createTree func(str string, _node *node)
	createTree = func(str string, _node *node) {
		for i := 0; i < len(str); i++ {
			__node := node{
				ch:       string(str[i]),
				children: nil,
			}
			_node.children = append(_node.children, &__node)
			cStr := ""
			for j, _ch := range str {
				if j != i {
					cStr += string(_ch)
				}
			}
			createTree(cStr, &__node)
		}
	}

	//create Tree
	var Tree []*node
	for i := 0; i < len(_string); i++ {
		_node := node{
			ch:       string(_string[i]),
			children: nil,
		}
		Tree = append(Tree, &_node)
		cStr := ""
		for j, _ch := range _string {
			if j != i {
				cStr += string(_ch)
			}
		}
		createTree(cStr, &_node)
	}

	var Res []string
	var CurrString string
	//build string
	var BuildString func(node *node)
	BuildString = func(node *node) {
		CurrString += node.ch
		if len(node.children) == 0 {
			Res = append(Res, CurrString)
		}
		for _, _node := range node.children {
			BuildString(_node)
		}
		CurrString = CurrString[:len(CurrString)-1]
	}
	for _, _node := range Tree {
		BuildString(_node)
	}

	//result
	for _, str := range Res {
		fmt.Println(str)
	}

}
