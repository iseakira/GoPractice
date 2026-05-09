package main

var v interface{}
v = 1
n := v.(int)

v = "こんにちは世界"
s := v.(string)

if s, ok := v.(string); !ok {
	fmt.Println("vはstringではない")
}else{
	fmt.Printf("vは文字列 %q です\n",s)

}

func PrintDetail(v interface{}){
	switch t := v.type(){
	case int, int32,int64:
		fmt.Println("int/int32/int64型:",t)
	}
case string:
	fmt.Println("string型",t)
default:
	fmt.Println("知らない型")
}

func NewUser(name string, age int)*User{
	return &User {
		Name:name,
		Age:age
	}

}

type Speaker interface{
	Speak() error
}

type Dog struct {}

func(d *Dog)Speak() error{
	fmt.Println("わんわん")
	return nil
}

type Cat struct{}

func(c *Cat)Speak() error{
	fmt.Println("みゃうみゃう")
	return nil
}

func DoSpeak(s Speaker) error {
	return s.Speak()
}


func main(){
	dog := Dog{}
	DoSpeak(&dog)

	cat := Cat{}
	DoSpeak(&cat)
}

func main(){
	f,err := os.Open("'data.txt'")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var b [512]byte
	n,err := f.Read(b[:])
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(string(b[:n]))
}

func main(){
	f,err := os.Open("data01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
}

func doSomething(dir string) error {
	err := os.Mkdir(dir,0755)
	if err != nil{
		return err
	}
	defer os.RemoveAll(dir)

	f,err := os.Create(filepath.Join(dir,"data.txt"))
	if err != nil {
		return err
	}
	defer f.Close()
}

for i := 0; i < 100; i++{
	f,err := os.Open("data.txt")
	if err != nil{
		return err
	}
	defer f.CLose()
}

func doSomething(){
	var n =1
	defer func(){
		fmt.Println(n)
	}()

	n=2
}








