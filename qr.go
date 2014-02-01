package main

import "math/rand"
import "fmt"
import "sort"

var root *node
type data [][2]int
type data1 [][2]int

type node struct {
	val        [2]int
	dimension   int
	left, right *node
}

func (a data) Len() int           { return len(a) }
func (a data) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a data) Less(i, j int) bool { return a[i][0] < a[j][0] }

func (a data1) Len() int           { return len(a) }
func (a data1) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a data1) Less(i, j int) bool { return a[i][1] < a[j][1] }

func buildTree(s [][2]int, l int, n **node) {
    if len(s) == 0 {
        return
    }
    if (l%2 == 0) {
        sort.Sort(data(s))
    } else {
        sort.Sort(data1(s))
    }
    median := (len(s))/2
    *n = &node{}
    (*n).val = s[median]
    (*n).dimension = l%2
    tmp := *n
    if median == 0 {
        return
    }
    buildTree(s [0:median], l+1, &tmp.left)
    buildTree(s [median+1:], l+1, &tmp.right)
}

func insert(s [][2]int) {
	cmd := 0
	for _, v := range s {
		tmp := &root
		for *tmp != nil {
			if (*tmp).val[cmd%2] > v[cmd%2] {
				tmp = &((*tmp).left)
			} else {
				tmp = &((*tmp).right)
			}
			cmd++
		}
		*tmp = &node{}
		(*tmp).val = v
		(*tmp).dimension = cmd % 2
	}
}

func trans(n *node, i int) {
	if n == nil {
		return
	}
	fmt.Println(n.val[0], n.val[1])
    fmt.Println(i, "my left")
	trans(n.left, i+1)
    fmt.Println(i, "my right")
	trans(n.right, i+1)
}

type stack struct {
	st    []*node
	index int
}

func (p *stack) push(v *node) {
    fmt.Println("push", v)
	p.st[p.index] = v
	p.index++
}

func (p *stack) pop() *node {
	if p.index == 0 {
		return nil
	}
	p.index--
	return p.st[p.index]
}

func (p *stack) empty() bool {
	if p.index == 0 {
		return true
	}
	return false
}

func dist(n *node, p [2]int) int {
	return sqr(n.val[0]-p[0]) + sqr(n.val[1]-p[1])
}

func sqr(a int) int {
	return a * a
}

func distAxis(n *node, p [2]int) int {
	if n.dimension == 0 {
		return sqr(n.val[0] - p[0])
	}
	return sqr(n.val[1] - p[1])
}

func (p *stack) top() *node {
	if p.index == 0 {
		return nil
	}
	return p.st[p.index-1]
}

func search(p [2]int) *node {
	if root == nil {
		return nil
	}
	sd := stack{}
	sd.st = make([]*node, 10)
	sd.index = 0
	curr_dist := 10000000
	var n *node = nil
    var tmp *node = root
	for {
		for {
			if tmp == nil {
				break
			}
			current := tmp.dimension
			if tmp.val[current] > p[current] || tmp.right == nil {
			        sd.push(tmp)
			        tmp = tmp.left
			} else {
					sd.push(tmp)
					tmp = tmp.right
			}
		}
		for {
			if sd.empty() == true {
				return n
			}
			parent := sd.pop()
			if parent == nil {
				return n
			}
			if val := distAxis(parent, p); val < curr_dist {
                if d := dist(parent, p); d < curr_dist {
                    n = parent
                    curr_dist = d
                }
                if parent.left == tmp {
				    tmp = parent.right
				    break
                } else {
				    tmp = parent.left
				    break
                }
			} else {
                tmp = parent
            }
		}
	}
}

func bruteForce(s [][2]int, p [2]int) [2]int {
    max := 10000000
    in := [2]int{}
    for _,val := range s {
        if d := did(val, p);d < max {
            max = d
            in = val
        }
    }
    return in
}

func did(m, n [2]int) int {
	return sqr(m[0]-n[0]) + sqr(m[1]-n[1])
}

func main() {
	s := [][2]int{}
    mynode := [2]int{12, 5}
	for i := 0; i < 10; i++ {
		x := rand.Intn(30)
		y := rand.Intn(30)
		s = append(s, [2]int{x, y})
	}
	//insert(s)
	buildTree(s, 0, &root)
	out := search(mynode)
    fmt.Println("brute force ", bruteForce(s, mynode))
	if out != nil {
		fmt.Println(out.val == bruteForce(s, mynode))
	}
}
