type Pair struct {
    num int
	fr  int
}

func topKFrequent(nums []int, k int) []int {
	var freq = make(map[int]int)

	for i := 0; i < len(nums); i++ {
		freq[nums[i]] += 1
	}

	var counter = make([]Pair, len(freq))
	for num, count := range freq {
		counter = append(counter, Pair{num: num, fr: count})
	}

	sort.Slice(counter, func(i, j int) bool {
		return counter[i].fr > counter[j].fr
	})

	var ans = make([]int, k)
	for i := 0; i < k; i++ {
		ans[i] = counter[i].num
	}

	return ans
}