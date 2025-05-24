package main

func print1toN(i,n int) {
	if i > n{
		return
	}

	println(i)
	print1toN(i+1,n)
}