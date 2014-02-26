package main

import "fmt"

func test(str string, findnth int) string {
    cha := []uint8{}
    for findnth > 0 {
        cha = append(cha, str[findnth%len(str)])
        findnth = findnth/len(str)
    }
    fmt.Println(string(cha))
    return string(cha)
}


func main() {
    arr := []uint8{}
    chars := "ADEFHNOPSUVY"
    tot := len(chars)
    findnth := 3365973428406169086
    findnth = findnth-1
    j:=tot
    for i:=tot;i<findnth;i=i*tot{
        findnth -= i
        j = i
    }
    for j > 0 {
        index := findnth/j
        arr = append(arr, chars[index])
        findnth=findnth-j*index
        j = j/tot
    }
    fmt.Println(test(chars, findnth))
}

