package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for index1, result1 := range nums {
		for index2, result2 := range nums {
			if index1 != index2 && result1+result2 == target {
				return []int{index1, index2}
			}
		}
	}
	return []int{}
}

// twoSumHashTable 使用 map 映射求解, 时间复杂度 O(n)
func twoSumHashTable(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}

func main() {

	fmt.Println(twoSumHashTable([]int{2, 5, 7, 9}, 7))
	fmt.Println(twoSum([]int{2, 5, 7, 9}, 7))

}
