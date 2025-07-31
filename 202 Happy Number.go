func isHappy(n int) bool {
    seen := make(map[int]bool)
    sumOfSquares := func (n int) int {
        sum := 0
        for n > 0 {
            d := n % 10
            sum += d * d
            n /= 10
        }
        return sum
    }
    for {
        if n == 1{
            return true
        }
        if seen[n] {
            return false
        }
        seen[n] = true
        n = sumOfSquares(n)
    }
}
