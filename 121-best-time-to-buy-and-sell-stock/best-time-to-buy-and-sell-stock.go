func maxProfit(prices []int) int {
    mini,maxProfit := prices[0],0
    n := len(prices)
    for i := 0 ; i < n ; i++{
        cost := prices[i] - mini
        maxProfit = max(maxProfit,cost)
        mini = min(mini,prices[i])
    }
    return maxProfit
}