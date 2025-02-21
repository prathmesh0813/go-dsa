func maxArea(height []int) int {
     left, right :=0, len(height)-1
    maxarea := 0

    for left < right{
        w := right - left
        h := min (height[right], height[left])
        a := w * h

        if a > maxarea{
            maxarea = a
        }

        if height[left] < height[right]{
            left++
        }else{
            right--
        }
    }
    return maxarea
}