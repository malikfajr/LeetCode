func twoSum(nums []int, target int) []int {
    var dataMap = make(map[int]int)

    for currIndex, currValue := range(nums) {
        if prevIndex, isPresent := dataMap[target-currValue]; isPresent {
            return []int{prevIndex, currIndex}
        }

        dataMap[currValue] = currIndex
    }

    return nil
}