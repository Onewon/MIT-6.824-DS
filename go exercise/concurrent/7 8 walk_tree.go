package main

// type Tree struct {
//     Left  *Tree
//     Value int
//     Right *Tree
// }

import (
	"golang.org/x/tour/tree"
	"fmt"
)
func WalkRecursive(t *tree.Tree, ch chan int) {
    if t != nil {
        WalkRecursive(t.Left, ch)
        ch <- t.Value
        WalkRecursive(t.Right, ch)
    }
}

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int){
	WalkRecursive(t, ch)
	close(ch)
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool{
	ch1, ch2 := make(chan int), make(chan int)
    go Walk(t1, ch1)
    go Walk(t2, ch2)
    for {
        n1, ok1 := <- ch1
        n2, ok2 := <- ch2
        if ok1 != ok2 || n1 != n2 {
        	return false
        }
        if !ok1 {
        	break;
        }
    }
    return true
} 
//return bool

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i :=range ch {
		fmt.Println(i)
	}
 	fmt.Println(Same(tree.New(1), tree.New(1))) //true
	fmt.Println(Same(tree.New(1), tree.New(2))) //false
}
