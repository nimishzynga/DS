package main

import "testing"
import (
    "fmt"
	"math/rand"
)

func Test_handleDown(t *testing.T) {
	s := [][2]int{}
	for i := 0; i < 20; i++ {
		x := rand.Intn(30)
		y := rand.Intn(30)
		s = append(s, [2]int{x, y})
	}
	//insert(s)
	buildTree(s, 0, &root)
    for j:=0;j<10;j++ {
	mynode := [2]int{}
	mynode[0] = rand.Intn(30)
    mynode[1] = rand.Intn(30)
	out := search(mynode)
    if out == nil || out.val != bruteForce(s, mynode) {
       fmt.Println("points",s)
       fmt.Println("brute",bruteForce(s, mynode),"myval",out.val,"point",mynode)
       trans(root, 0)
        t.Fail()
    }
    }
}
