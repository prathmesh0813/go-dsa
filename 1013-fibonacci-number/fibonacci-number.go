func fib(n int) int {
    // if n == 0{
    //     return 0
    // }
    // if n == 1{
    //     return 1
    // }

    // n1,n2 := 0,1
    // for i:=2;i<=n;i++{
    //     n1,n2 = n2,n1+n2
    // }
    // return  n2

    //Using reursion 
    if n <=1 {
        return n
    }

    return  fib(n-1) +  fib(n-2)
}