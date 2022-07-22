package main

import "math/rand"

type RandomizedSet struct {
	values []int
	valueIndexMap map[int]int
}


func RandomizedSetConstructor() RandomizedSet {
	return RandomizedSet{[]int{}, map[int]int{}}
}


func (rs *RandomizedSet) Insert(val int) bool {
	if _, exists := rs.valueIndexMap[val]; exists {
		return false
	}
	rs.valueIndexMap[val] = len(rs.values)
	rs.values = append(rs.values, val)
	return true
}


func (rs *RandomizedSet) Remove(val int) bool {
	index, ok := rs.valueIndexMap[val]
	if !ok {
		return false
	}
	last := len(rs.values) - 1
	rs.valueIndexMap[rs.values[index]] = index
	rs.values[index] = rs.values[last]
	rs.values = rs.values[:last]
	delete(rs.valueIndexMap, val)
	return true
}

func (rs *RandomizedSet) GetRandom() int {
	return rs.values[rand.Intn(len(rs.values))]
}