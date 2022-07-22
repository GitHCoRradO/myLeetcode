package main

func reverseString(s []byte)  {
	low := 0
	up := len(s) - 1
	for low < up {
		lowC := s[low]
		s[low] = s[up]
		s[up] = lowC
		low++
		up--
	}
}
