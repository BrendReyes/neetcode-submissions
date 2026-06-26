func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }

    left := 0 
    maxProj := 0 

    for right := 1; right < len(prices); right++ {
        if prices[right] < prices[left] {
            left = right
        } else {
			currProf := prices[right] - prices[left]
			if currProf > maxProj {
				maxProj = currProf
			}
        }
    }

    return maxProj
}