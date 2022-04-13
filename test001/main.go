package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
)

type person struct {
	age int
	name string
}

func main()  {
	var handler http.Handler
	http.ListenAndServe("127.0.0.1:80", handler)
	fmt.Println("1111111")
	fmt.Println("输入的参数：", os.Args[1])
/*
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
*/


/*--------------array-------------*/
	fmt.Printf("/*--------------array-------------*/\n")
	var array01[5] int
	array02 := [5]int{0,1,2,3,4}
	array03 := [5]int{0:8, 1:8}
	fmt.Printf("%d\n", array01[3])
	fmt.Printf("%d\n", array02[0])
	fmt.Printf("%d\n", array03[0])

	for i, v:=range array03 {
		fmt.Printf("索引[%d], 值[%d]\n", i, v)
	}

	array04 := [4]*int{0:new(int), 3:new(int)}
	*array04[0] = 1
	//*array04[1] = 1	// error
	array04[1] = new(int)
	*array04[1] = 9
	fmt.Printf("%d-%d\n", array04[1], *array04[1])

	array05 := [5]int{1:2, 3:4}
	fmt.Println("before:", array05)
	modify_array0(&array05)
	fmt.Println("after :", array05)

	var array06 [6]int
	fmt.Println("before:", array06)
	modify_array1(array06)
	fmt.Println("after :", array06)


/*--------------slice-------------*/
	fmt.Printf("/*--------------slice-------------*/\n")
	slice := []int{0,1,2,3,4}
	for _, v:=range slice {
		// 不想要索引，可以使用_来忽略它
		fmt.Println(v)
	}

	// 值的方式传递，占用内存小
	fmt.Println("before modify_slice:", slice)
	fmt.Printf("before modify_slice addr:%p\n", &slice)
	modify_slice(slice)
	fmt.Println("after modify_slice :", slice)

	slice1 := slice[:]
	slice2 := slice[0:]
	slice3 := slice[:4]
	slice4 := slice[1:3]
	fmt.Println(slice)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)	//左闭右开

	newSlice := slice[1:3]
	fmt.Printf("newSlice长度:%d,容量:%d\n",len(newSlice),cap(newSlice))
	newSlice = append(newSlice, 10, 20, 30)
	newSlice = append(newSlice, slice...)	//slice追加到newSlice
	fmt.Println(newSlice)

	fmt.Printf("/*--------------slice-------------*/\n")
/*--------------map-------------*/
	fmt.Printf("/*--------------map-------------*/\n")
	dict := make(map[string] int)
	dict["cccc"] = 7
	dict["aaaa"] = 9
	dict["bbbb"] = 8

	sort_map(dict)

	delete(dict, "aaaa")
	dict1 := map[string]int{}
	dict1["aaaa"] = 8

	fmt.Printf("/*--------------map-------------*/\n")

/*--------------struct-------------*/
	fmt.Printf("/*--------------struct-------------*/\n")
	//var p person
	p := person{12, "jim"}
	modify_struct_0(&p)
	fmt.Println(p)
	modify_struct_1(p)
	fmt.Println(p)

	fmt.Println(p.String_name())
	fmt.Println(p.Int_age())

	p.modify_value()
	fmt.Println("值接收者，修改无效：", p.String_name())
	p.modify_point()
	fmt.Println("指针接收者，修改有效：", p.String_name())
	fmt.Printf("/*--------------struct-------------*/\n")


/*--------------多值返回-------------*/
	openfile()
	fmt.Println(add(1,2))


/*--------------可变参数-------------*/
	print("qqq", "ww", "e")
}

func modify_array0(a *[5]int) {
	//传递数组指针，修改原数组
	a[1] = 3
	fmt.Println("modify_0:", *a)
}

func modify_array1(a [6]int) {
	//复制数组，不修改原数组
	a[1] = 3
	fmt.Println("modify_1:", a)

}

func modify_slice(slice []int) {
	fmt.Printf("%p\n", &slice)	//两个slice地址不一样
	slice[1] = 10
}

func sort_map(dict map[string]int) {
	fmt.Println("Here is sort_map.")
	var names []string
	for name := range dict {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, key:= range names {
		fmt.Println(key, dict[key])
	}
}

func modify_struct_0(p *person) {
	//通过传递结构体指针，可以修改age的值
	p.age = p.age + 10
}

func modify_struct_1(p person) {
	// 结构体传递的是其本身以及里面的值的拷贝，无法修改age的值
	p.age = p.age + 10
}

func (p person) String_name() string{
	return "the person name is " + p.name
}

func (p person) Int_age() int{
	return p.age
}

func (p person) modify_value() {
	p.name = "sunlan"
}

func (p *person) modify_point() {
	p.name = "sunlan"
}

func openfile() {
	file, err := os.Open("/Users/apple/go_workspace/test001")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(file.Chdir())
}

func add(a, b int) (int, error) {
	return a+b, nil
}

func print(input ...interface{}) {
	for _, v:=range input {
		fmt.Print(v)
	}
	fmt.Println("\n", input)
}