package main

import "fmt"

type Student struct {
	ID              int
	Name            string
	IsActiveStudent bool
	Score           float32
	Address         []string
	Dept
}

type Dept struct {
	DeptName string
	DeptId   int
	Address  []string
}

func main() {

	// d := Dept{11, "IT"}
	sd := new(Dept)
	sd = &Dept{DeptName: "Civil Dept", DeptId: 21, Address: []string{}}
	newList := new([5]int)
	cplx := 5.5
	fmt.Println("", cplx)
	fmt.Println("New List ", newList)
	fmt.Println("USING NEW ", sd.DeptName)
	dd := Dept{DeptId: 11, DeptName: "CSE", Address: []string{"Main Road Building", "Entrance Hall Side"}}

	stud1 := Student{1, "ABCD", false, 3.3, []string{
		"A",
		"B",
		"C",
	}, dd}
	fmt.Println("start", dd)

	fmt.Printf("%+v\n", stud1)

	fmt.Println("Student Address", stud1.Address)
	fmt.Println("Student Dept Name", stud1.DeptName)
	fmt.Println("Student Dept Name ", stud1.Dept.DeptName)
	fmt.Println("Student Dept Address", stud1.Dept.Address)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	TestName(11)

	func() {
		fmt.Println("Self executing function ")
	}()

	fmt.Println("BEFORE ", stud1.Name)
	UpdateStud(&stud1)
	fmt.Println("After ", stud1.Name)
}

func UpdateStud(s *Student) {
	s.Name = "XYZZZZ"
}
func TestName(Id int) {
	fmt.Println("SECRET FUNCTIONSS", Id)
}

