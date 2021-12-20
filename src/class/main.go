package main

import "fmt"

// 学生信息
type student struct {
	id   int
	name string
	age  int
}

// 信息系统
type information struct {
	StudentList map[int]*student
}

// 构造函数
func newStudent(id, age int, name string) *student {
	return &student{
		id:   id,
		age:  age,
		name: name,
	}
}

func newInformation() *information {
	return &information{
		StudentList: map[int]*student{},
	}
}

func (i *information) addStudent(student *student) {
	i.StudentList[student.id] = student
}

func (i *information) getStudent(id int) *student {
	return i.StudentList[id]
}

func (i *information) getStudents() {
	for k, v := range i.StudentList {
		fmt.Println(k, *v)
	}
}

func showMenus() {
	fmt.Println("欢迎来到学员管理系统")
	fmt.Println("请选择你需要操作的菜单")
	fmt.Println("1.添加学员")
	fmt.Println("2.编辑学员")
	fmt.Println("3.查询学员")
	fmt.Println("4.退出")
}

func main() {
	flag := true
	input := 0
	c1 := newInformation()
	id := 1
	showMenus()
	for flag {
		fmt.Println("请输入选择的菜单")
		fmt.Scanf("%d", &input)
		switch input {
		case 1:
			s1 := newStudent(id, 20, "luyuan")
			c1.addStudent(s1)
			id++
		case 2:
			innerId := 0
			fmt.Println("请输入修改的学员id")
			fmt.Scanf("%d", &innerId)
			getStudent := c1.getStudent(innerId)
			if getStudent == nil {
				fmt.Println("学员不存在")
				continue
			}
			getStudent.name = "已修改"
		case 3:
			c1.getStudents()
		case 4:
			flag = !flag
		default:
			fmt.Println("菜单输入错误")
		}
	}
}

func Test() {

}
