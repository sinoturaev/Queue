package main

import "fmt"

type List struct {
	root *ListNode // first element
	end *ListNode // last element
	length int // dlinna
}

func (receiver *List) PrintList()  {
	fmt.Println("Start print...")
	current := receiver.root
	for current != nil {
		fmt.Println(current)
		current = current.Next
	}
	fmt.Println("Finish")
}

func (recevier *List) len() int  {
	return recevier.length
}

func (recevier *List) Add(node ListNode)  {
	if recevier.len() == 0 {
		recevier.root = &node
		recevier.end = &node
		recevier.length++
		return
	}
	////
	LastElement := recevier.end
	recevier.end.Next = &node
	recevier.end = &node
	recevier.end.Prev = LastElement
	recevier.length++
}

/*func (recevier *List) popBack(node ListNode)  {
	if recevier.len() ==
}*/

type ListNode struct {
	Prev *ListNode
	Next *ListNode
	Name string
	Purchases int



}



func main() {
	person := ListNode{
		Prev:      nil,
		Next:      nil,
		Name:      "Sino",
		Purchases: 150,
	}

	queue := List{}
	queue.Add(person)
	fmt.Println(queue)


	person = ListNode{
		Prev:      nil,
		Next:      nil,
		Name:      "Azam",
		Purchases: 110,
	}
    queue.Add(person)
	fmt.Println(queue)

	person = ListNode{
		Prev:      nil,
		Next:      nil,
		Name:      "Raboni",
		Purchases: 120,
	}
	queue.Add(person)
	fmt.Println(queue)

	queue.PrintList()

	fmt.Println(queue.root)
	fmt.Println(queue.end)


}

