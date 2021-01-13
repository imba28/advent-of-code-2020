package main

import (
	"container/list"
	"fmt"
)

func List(cards []int) *list.List {
	l := list.New()
	for i := range cards {
		l.PushBack(cards[i])
	}
	return l
}

func Play(a, b *list.List) *list.List {
	for a.Len() > 0 && b.Len() > 0 {
		topA := a.Front().Value.(int)
		topB := b.Front().Value.(int)

		a.Remove(a.Front())
		b.Remove(b.Front())

		if topA > topB {
			a.PushBack(topA)
			a.PushBack(topB)
		} else {
			b.PushBack(topB)
			b.PushBack(topA)
		}
	}

	if a.Len() == 0 {
		return b
	}
	return a
}

func score(l *list.List) int {
	v := 0

	for i := 1; l.Len() > 0; i++ {
		front := l.Back()
		l.Remove(front)

		v += i * front.Value.(int)
	}

	return v
}

func main() {
	cardsA := List([]int{18, 50, 9, 4, 25, 37, 39, 40, 29, 6, 41, 28, 3, 11, 31, 8, 1, 38, 33, 30, 42, 15, 26, 36, 43})
	cardsB := List([]int{32, 44, 19, 47, 12, 48, 14, 2, 13, 10, 35, 45, 34, 7, 5, 17, 46, 21, 24, 49, 16, 22, 20, 27, 23})

	//cardsA := List([]int{9, 2, 6, 3, 1})
	//cardsB := List([]int{5, 8, 4, 7, 10})

	winner := Play(cardsA, cardsB)
	fmt.Println(score(winner))
}
