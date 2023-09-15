func removeDuplicates(nums []int) int {
    p1 := 0
    p2 := 1

    for i := 1; i < len(nums); i++ {
        if nums[p1] != nums[i] {
            nums[p2] = nums[i]
            p1++
            p2++
        }
    }

    return p2;
}