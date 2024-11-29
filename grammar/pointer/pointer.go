package main

import (
	"encoding/json"
	"fmt"
)

/*
指针
*/
func main() {
	basic()

	//	结构体
	structFn()

	// 结构体嵌套
	nesting()

	// 结构体数据转换
	jsonChange()

	// 复杂数据转换
	testJson()
	testStruct()
}

type Person struct {
	name   string
	age    int
	sex    string
	height int
	hobby  []string
	map1   map[string]string
	// /* address address */
	// 嵌套一个匿名结构体
	address
	email
}

type address struct {
	name  string
	phone string
	city  string
}

type email struct {
	city string
}

func basic() {
	var a = 10 // 10
	var b = &a // 指针地址
	var c = *b // 10
	println(a, b, c)
}

/*
结构体自定义方法
*/
func (p *Person) printInfo() {
	fmt.Println(" ")
	fmt.Printf("结构体自定义方法 ---> 姓名：%v 年龄：%v\n", p.name, p.age)
}

/*
结构体中修改值类型数据
*/
func (p *Person) setInfo(name string, age int) {
	p.name = name
	p.age = age
}

/*
结构体
*/
func structFn() {
	fmt.Println(" ")
	var p1 Person // 实例化 Person 结构体
	p1.name = "结构体名称"
	p1.age = 18
	p1.sex = "man"
	fmt.Printf("结构体数据：%#v 类型 %T\n", p1, p1)

	// 结构体中自定义的方法
	p1.printInfo() // 结构体名称 18

	// 结构体中修改值类型的数据 使用指针
	p1.setInfo("李四修改值类型数据", 100)

	p1.printInfo() // 李四修改值类型数据 100

	// 结构体中添加切片数据
	p1.hobby = make([]string, 3, 3)
	p1.hobby = append(p1.hobby, "吃饭")
	p1.hobby = append(p1.hobby, "睡觉")
	p1.hobby = append(p1.hobby, "写代码")

	fmt.Println(" ")
	fmt.Printf("p1 ---> %#v \n", p1)
}

/*
结构体嵌套
*/
func nesting() {
	fmt.Println(" ")
	p2 := Person{
		name: "嵌套结构",
	}
	p2.address = address{
		name:  "张三的地址",
		phone: "12345678910",
		city:  "未知",
	}
	// 会在 Person 结构体里面查找 city，查询不到再往嵌套结构体里面查找
	//p2.city = "上海"
	fmt.Printf("p1 ---> %#v \n", p2)
}

type Student struct {
	Id     int    `json:"id"` // 结构体标签
	Gender string `json:"gender"`
	Name   string
	sno    string // 首字母小写为私有字段不能被json包访问
}

/*
json 结构体转换
*/
func jsonChange() {
	fmt.Println(" ")
	s := Student{
		Id:     110,
		Gender: "男",
		Name:   "李四",
		sno:    "S001",
	}
	fmt.Printf("s ---> %#v \n", s)

	// 转成JSON
	jsonS, _ := json.Marshal(s)
	jsonStr := string(jsonS)
	fmt.Printf("s 转成JSON的值： ---> %v \n", jsonStr)

	fmt.Println(" ")
	// JSON 转 struct结构体
	var str = ` {"Id":110,"Gender":"男","Name":"李四","Sno":"S001"}`
	var sJson Student
	err := json.Unmarshal([]byte(str), &sJson) // 接收的变量需要取指针地址
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Printf("JSONs转 struct 的值： ---> %v \n", jsonStr)
}

type Class struct {
	Title   string    `json:"title"`
	Student []Student `json:"student"`
}

func testJson() {
	fmt.Println(" ")
	tStruct := Class{
		Title:   "测试班级",
		Student: make([]Student, 0),
	}
	for i := 1; i < 20; i++ {
		s := Student{
			Id:     i,
			Name:   fmt.Sprintf("stu_%v", i),
			Gender: "男",
		}
		tStruct.Student = append(tStruct.Student, s)
	}

	jsonByte, err := json.Marshal(tStruct)
	if err != nil {
		fmt.Println("错误", err)
	} else {
		jsonStr := string(jsonByte)
		fmt.Printf("JSON 的值： ---> %v \n", jsonStr)
	}
}

func testStruct() {
	fmt.Println(" ")
	str := `{"title":"测试班级","student":[{"id":1,"gender":"男","Name":"stu_1"},{"id":2,"gender":"男","Name":"stu_2"},{"id":3,"gender":"男","Name":"stu_3"},{"id":4,"gender":"男","Name":"stu_4"},{"id":5,"gender":"男","Name":"stu_5"},{"id":6,"gender":"男","Name":"stu_6"},{"id":7,"gender":"男","Name":"stu_7"},{"id":8,"gender":"男","Name":"stu_8"},{"id":9,"gender":"男","Name":"stu_9"},{"id":10,"gender":"男","Name":"stu_10"},{"id":11,"gender":"男","Name":"stu_11"},{"id":12,"gender":"男","Name":"stu_12"},{"id":13,"gender":"男","Name":"stu_13"},{"id":14,"gender":"男","Name":"stu_14"},{"id":15,"gender":"男","Name":"stu_15"},{"id":16,"gender":"男","Name":"stu_16"},{"id":17,"gender":"男","Name":"stu_17"},{"id":18,"gender":"男","Name":"stu_18"},{"id":19,"gender":"男","Name":"stu_19"}]}`
	c1 := &Class{}
	s1 := json.Unmarshal([]byte(str), c1)
	if s1 != nil {
		fmt.Println("json err", s1)
		return
	}
	fmt.Printf("struct 的值： ---> %v \n", c1)
}
