func findDisappearedNumbers(nums []int) []int {
    for _, num := range nums {
        index := abs(num) - 1
        if nums[index] > 0 {
            nums[index] = -abs(nums[index])
        }
    }
    result := []int{}
    for i := 0; i < len(nums); i++ {
        if nums[i] > 0 {
            result = append(result, i+1)
        }
    }
    return result

}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}