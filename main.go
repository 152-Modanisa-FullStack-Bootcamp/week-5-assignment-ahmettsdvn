package main

import (
	"fmt"
	"sort"
    "strings"

	s "bootcamp/service"
)

func SortString(w string) string {
    s := strings.Split(w, "")
    sort.Strings(s)
    return strings.Join(s, "")
}

func main() {
	fmt.Println("Bootcamp week 5 assignment!")
	words := "apple,bat,cat,goodbye,hello,yellow,why"
	//sum, overflow :=s.AddUint32(1, 1)
	//v := s.VariadicSet([]interface{}{1,uint32(1),"first",2,uint32(2),"second",1,uint32(2),"first"})
    fmt.Println("f => ",s.WordsSplit([2]string{"derebere",words}))
}
// 4294967290
// 2147483647
