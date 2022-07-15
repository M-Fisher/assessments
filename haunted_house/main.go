package main

func main() {
	return
}

func getMaxGroup(n int, people [][]int) int {
	// list of possible group sizes and number of people ready to go with this number of other people
	groups := map[int]int{}
	for _, rules := range people {
		for i := rules[0]; i <= rules[1]; i++ {
			groups[i]++
		}
	}
	max := 0
	for i, total := range groups {
		if i+1 <= n { // If group size + person himself less than total number of people

			// number of people ready to go with i other can be more than i ->
			// not everyone can go with this group
			readyToGo := min(total, i+1)
			if readyToGo > max {
				max = readyToGo
			}
		}
	}

	return max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
