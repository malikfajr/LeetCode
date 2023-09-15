func removeDuplicates(nums []int) int {
    p1 := 0

    for i := 1; i < len(nums); i++ {
        if nums[p1] != nums[i] {
            p1++
            nums[p1] = nums[i]
        }
    }

    return p1 + 1;
}