func containsDuplicate(nums []int) bool {
    var dataMap = make(map[int]bool)

    for i := 0; i < len(nums); i++ {
      if dataMap[nums[i]] {
        return true
      }

      dataMap[nums[i]] = true;
    }

    return false
}