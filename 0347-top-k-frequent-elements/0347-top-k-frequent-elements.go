func topKFrequent(nums []int, k int) []int {
    var counts = make(map[int]int)
    
    for i := 0; i < len(nums); i++ {
        counts[nums[i]] += 1;
    }
    
    var result = []int{}
	for k > 0 {
		i := -1
		var v int
		for num, count := range counts {
			if i < count {
			i = count
			v = num
			}
		
		}
		result = append(result, v)
		delete(counts, v)
		k--;
	}
	return result
}