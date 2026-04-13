func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        sybau := target - nums[i]
        if _, ok := m[sybau]; ok {
            return []int{m[sybau], i}
        }
        m[nums[i]] = i
    }
    return nil
}
