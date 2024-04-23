package main

import "fmt"

type Queue struct {
	Array []int
}
type Stack struct {
	Array []int
}
type Car interface {
	Speed(int) int
	Fly(string) string
}

func main() {

	fmt.Println("Start")
	que := Queue{}
	fmt.Println("Added element : ", que.Enque(10))
	fmt.Printf("Queue %+v", que)
	fmt.Println("Added element : ", que.Enque(20))
	fmt.Printf("Queue %+v", que)
	fmt.Println("Added element : ", que.Enque(30))
	fmt.Printf("Queue %+v", que)
	fmt.Println("Added element : ", que.Enque(40))
	fmt.Printf("Queue %+v", que)
	fmt.Println("Added element : ", que.Enque(50))
	fmt.Println("Added element : ", que.Enque(60))
	fmt.Println("Added element : ", que.Enque(70))
	fmt.Println("Added element : ", que.Enque(80))
	//L I F O
	//F I L O

	//Que
	// F I F O
	// L I L O
	fmt.Printf("Queue %+v", que)
	fmt.Println("Dequeue ", que.Dequeue())
	fmt.Printf("Queue %+v", que)
	fmt.Println("Dequeue ", que.Dequeue())
	fmt.Printf("Queue %+v", que)
	fmt.Println("Dequeue ", que.Dequeue())
	fmt.Printf("Queue %+v", que)
	fmt.Println("Dequeue ", que.Dequeue())
	fmt.Printf("Queue %+v", que)
	fmt.Println("Dequeue ", que.Dequeue())
	fmt.Printf("Queue %+v", que)
	fmt.Println("Dequeue ", que.Dequeue())
	fmt.Printf("Queue %+v", que)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	nQ := Queue{}
	fmt.Println("", nQ.Enque(22), nQ)

	nS := Stack{}
	fmt.Println("", nS.Enque(33), nS)

}

func (s *Queue) Enque(ele int) int {
	fmt.Println("QUEUE IS CALLED")
	s.Array = append(s.Array, ele)
	return ele
}

func (s *Stack) Enque(ele int) int {
	fmt.Println("STACK IS CALLED")
	s.Array = append(s.Array, ele)
	return ele
}

func (s *Queue) Dequeue() int {
	firstAddedElement := s.Array[0]
	s.Array = s.Array[1:]
	return firstAddedElement
}

