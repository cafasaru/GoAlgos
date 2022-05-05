package numinlist

// NumInSlice takes a
func NumInSlice(nums []int, find int) bool {
	if len(nums) == 0 {
		return false
	}
	for _, num := range nums {
		if num == find {
			return true
		}
	}

	return false
}
