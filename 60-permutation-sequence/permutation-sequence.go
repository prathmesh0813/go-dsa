func getPermutation(n int, k int) string {
    fact := 1
    numbers := []int{}
    for i:=1;i<n;i++{
        fact *=i
        numbers = append(numbers,i)
    }
    numbers = append(numbers,n) 

    k = k-1
    ans := ""
    for {
        index := k/fact
        ans += strconv.Itoa(numbers[index])

        numbers = append(numbers[:index],numbers[index+1:]...)

        if len(numbers) == 0{
            break
        }

        k= k%fact
        fact = fact/len(numbers)
    }
    return ans

}