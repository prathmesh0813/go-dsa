package main



// Given an integer array nums. Return the number of inversions in the array.
// Two elements a[i] and a[j] form an inversion if a[i] > a[j] and i < j.
// It indicates how close an array is to being sorted.
// A sorted array has an inversion count of 0.
// An array sorted in descending order has maximum inversion.
// Examples:
// Input: nums = [2, 3, 7, 1, 3, 5]

// Output: 5

// Explanation: The responsible indexes are:

// nums[0], nums[3], values: 2 > 1 & indexes: 0 < 3

// nums[1], nums[3], values: 3 > 1 & indexes: 1 < 3

// nums[2], nums[3], values: 7 > 1 & indexes: 2 < 3

// nums[2], nums[4], values: 7 > 3 & indexes: 2 < 4

// nums[2], nums[5], values: 7 > 5 & indexes: 2 < 5

func merge(arr [] int,low,mid,high int) int {
	temp := make([]int,0,high-low+1)
	count := 0
	left := low
	right := mid + 1

	for left <= mid && right <= high{
		if arr[left] <= arr[right]{
			temp = append(temp, arr[left])
			left++
		}else{
			temp = append(temp, arr[right])
			count += (mid-left+1)
			right++
		}
	}


	for left <= mid{
		temp = append(temp, arr[left])
		left++
	}

	for right <= high{
		temp = append(temp, arr[right])
		right++
	}

	for i := low;i<=high;i++{
		arr[i] =temp[i-low]
	}
	return count
}

func mergeSort(arr[]int,low, high int) int{
	count := 0
	if low < high{
		mid := (low+high)/2
		count += mergeSort(arr,low,mid)
		count += mergeSort(arr,mid+1,high)
		count += merge(arr,low,mid,high)
	}
	return count
}

func numberOfInversions(arr []int,n int) int {
	return mergeSort(arr,0,n-1)
}