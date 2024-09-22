package main

func twoSum1(nums []int, target int) []int {
	hashTable := map[int]int{}

	for i, x := range nums {
		if v, ok := hashTable[target-x]; ok {
			return []int{i, v}
		}
		hashTable[x] = i
	}
	return nil
}

func main() {

}
