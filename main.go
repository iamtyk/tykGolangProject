package main

import (
	"fmt"
	"sort"
	"strings"
	"tykProject/MinHeap"
)

func main() {
	strings.Index("abc", "a")
	maximizeXor([]int{0, 1, 2, 3, 4}, [][]int{{3, 1}, {1, 3}, {5, 6}})

	fmt.Println("hello world")
	heap := MinHeap.Constructor(2, []int{1, 2, 3, 4})
	heap.Add(2)
	fmt.Println("testGitCommit")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carrat := 0
	result := &ListNode{}
	current := result
	for l1 != nil || l2 != nil {
		val := 0
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}

		val += carrat
		carrat = val / 10
		val = val - carrat*10
		if current == nil {
			current = &ListNode{}
			current.Val = val
		} else {
			current.Val = val
		}
		if carrat > 0 {
			current.Next = &ListNode{}
			current = current.Next
			current.Val = carrat
		}

		current = current.Next
	}

	if carrat > 10 {
		current.Next = &ListNode{}
		current = current.Next
		current.Val = carrat % 10

		current.Next = &ListNode{}
		current = current.Next
		current.Val = carrat / 10
	} else if carrat > 0 {
		current.Next = &ListNode{}
		current = current.Next
		current.Val = carrat
	}

	return result
}

type Tire struct {
	children [2]*Tire
}

func (this *Tire) insertTire(num int) {
	tail := this
	for i := num; i > 0; i = i >> 1 {
		index := i & 1
		if tail.children[index] == nil {
			tail.children[index] = &Tire{}
		}
		tail = tail.children[index]
	}
}

func (this *Tire) findMaxXoR(val int) (ans int) {
	tail := this
	i := val
	idx := 0
	for tail != nil || i > 0 {
		xor := i & 1
		if tail != nil {
			if tail.children[xor^1] != nil {
				ans = 1<<idx | ans
				xor ^= 1
			}
			tail = tail.children[xor]
		} else if i > 0 {
			bit := 1 << idx & i
			ans = bit | ans
		}
		idx++
		if i > 0 {
			i = val >> idx
		}
	}
	return
}

func maximizeXor(nums []int, queries [][]int) []int {
	sort.Ints(nums)
	for i := range queries {
		queries[i] = append(queries[i], i)
	}
	sort.Slice(queries, func(i, j int) bool {
		return queries[i][1] < queries[j][1]
	})
	ans := make([]int, len(queries))
	t := &Tire{}
	idx, n := 0, len(nums)
	for _, q := range queries {
		x, m, qid := q[0], q[1], q[2]
		for idx < n && nums[idx] <= m {
			t.insertTire(nums[idx])
			idx++
		}
		if idx == 0 {
			ans[qid] = -1
		} else {
			ans[qid] = t.findMaxXoR(x)
		}
	}
	return ans
}
