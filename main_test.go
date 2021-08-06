package main

import "testing"

func TestList_Add(t *testing.T) {
	q := List{}

	if q.length != 0 {
		want := 0
		got := q.len()
		t.Errorf("method len() in Method ( add ) is not corret for empty queue got %d want %d\n", got, want)
	}
	person := ListNode{
		Prev:      nil,
		Next:      nil,
		Name:      "Raboni",
		Purchases: 120,
	}
	q.Add(person)
	if q.length != 1 {
		want := 0
		got := q.len()
		t.Errorf("method len() in Method( add ) is not coorect for empty queue got %d want %d\n", got, want)
		// TODO: TEST for your new method getFirst Element and get getLast Element
	}
}