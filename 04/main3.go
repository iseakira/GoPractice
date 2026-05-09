package main

import "fmt"

type User struct {
	Name string
	Age int
}

type Value int

func (v *Value) Add(n Value) {
	*v += n
}



func main(){
	v := Value(1)
	v = v.Add(2)
}





func showName(user *User){
	fmt.Println(user.Name)
}

func main() {

	user := User{
	Name:"Bob",
	Age:18,
}
showName(&user)


}