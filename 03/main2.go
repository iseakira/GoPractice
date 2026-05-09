package main

if x ==2 && y == 3{
	doSomething
}

if user,err := userName(); err == nil {
	fmt.Println(user.Name)
}else {
	log.Println(err)
}

switch i {
case 1:
	fmt.Println("one")
case 2:
	fmt.println("two")
	fallthrough
case 3,4:
	fmt.Println("three")
default:
	fmt.Println("other")
}

switch {
case i > hit:
	fmt.Println("bigger than hit")
case i < hit:
	fmt.Println("less than hit")
case i == hit:
	fmt.Println("equal to hit")
}

for i := 0; i < lens(s); i++{
	//do something
}

for i < 3{

}

for k,v := range obj {

}

for  i, v := range arr {

}

var a [4] int
a[0] = 1

var a [] int


 a := make([]int, 3)
 a[0] = 1
 a[1] = 2
 a[2] = 3

 a := []int{1,2,3}

 arr1 := [2][3]int{
	{1,2,3},
	{4,5,6}
 }

 arr2 := [][]int{
	{1,2,3},
	{4,5,6}
 }

 a := []int{1,2,3}
 fmt.Println(a)

 a = append(a,4,5,6)
 fmt.Println(a)
 fmt.Println("aの長さは %d\n",len(a))

 a := make([]int,0,100)
 for i :=0; i < 100; i++ {
	a = append(a,i)
 }

 a2 := make([]int, 0, len(a/2))
 for i := 0; i < len(a); i ++ {
	if i % 2 == 0 {
		a2 = append(a2,a[i])
	}
 }
 a=a2

 s := "Hello"
 s += name

 b := []bytes(s)
 b[0] = "h"
 s = string(b)

 s := "今日は世界"
 rs := []rune(s)
 rs[4] = 'は'
 s=string(rs)

 var b [4]byte

 b2 := b[:]

var m map[string]int

m := make(map[string]int)
m["jone"]=21
m["Bob"]=18
m["Mark"]=33

m := make(map[string]int,10)

m := map[string]int{
	"John":21,
	"Bob":18,
	"Mark":33,
}

delete(m,"Bob")

for k,v := range m {


}

keys := []string{}

for k := range m {
	keys = appned(keys,k)
}

sort.Strings(keys)
for _, k := range keys {
	fmt.Println()
}

v, ok := m["zoo"]
if ok {
	fmt.Println(v)
}

