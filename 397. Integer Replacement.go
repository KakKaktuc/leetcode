func integerReplacement(n int) int {
    count := 0
    for n != 1 {
        if n % 2 == 0{
            n = n / 2
            count++
        } else if n == 3 {
            n = n - 1
            count++
        }else if n % 4 == 3{
            n = n + 1
            count++
        } else if n % 4 == 1 {
            n = n -1
            count++
        }
    }
    return count
}