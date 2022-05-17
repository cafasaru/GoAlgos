package sumlist

func sumList(list []int) int {
	total := 0
	for _, v := range list {
		total += v
	}

	return total
}
