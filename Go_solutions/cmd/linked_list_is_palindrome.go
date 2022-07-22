package main

func linkedListIsPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}
	if head.Next == nil {
		return true
	}
	slice := make([]int, 0)
	for head.Next != nil {
		slice = append(slice, head.Val)
		head = head.Next
	}
	slice = append(slice, head.Val)
	low := 0
	high := len(slice)-1
	for ; low <= high; {
		if slice[low] == slice[high] {
			low +=1
			high -= 1
			continue
		} else {
			return false
		}
	}
	return true
}
