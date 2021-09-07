package main

import "fmt"

// Person 自定义类型
type Person struct {
	Name string
	City string
	Age  int
}

// NewPerson 构造函数
func NewPerson(name, city string, age int) *Person {
	return &Person{
		Name: name,
		City: city,
		Age:  age,
	}
}

// 方法接受者

// Getter
func (p *Person) getName() string {
	return p.Name
}

func (p *Person) getCity() string {
	return p.City
}

func (p *Person) getAge() int {
	return p.Age
}

// Setter
func (p *Person) setName(name string) {
	p.Name = name
}

func (p *Person) setCity(city string) {
	p.City = city
}

func (p *Person) setAge(age int) {
	p.Age = age
}

// Worker 继承类型
type Worker struct {
	Level  string
	Job    string
	Person // 类型和参数名字形同可以使用匿名类型
}

func main() {
	p1 := NewPerson("陆远", "上海", 18)
	fmt.Println(p1)
	fmt.Println(p1.getAge())
	fmt.Println(p1.getCity())
	fmt.Println(p1.getName())
	p1.setName("测试")
	p1.setCity("北京")
	p1.setAge(30)
	fmt.Println(p1)

	worker := Worker{
		Level: "CEO",
		Job:   "打杂",
		Person: Person{
			Name: "老板",
			City: "上海",
			Age:  80,
		},
	}
	fmt.Println(worker)
	worker.setName("农名")
	worker.Level = "BOSS" // setter必须用指针（全部用指针）
	fmt.Println(worker)
}
