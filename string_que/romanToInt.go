package main

func RomanToInt(s string ) int {
	result :=0

	romanMap := map[byte]int{
		'I' : 1, 'V' : 5, 'X' : 10, 'L' : 50 , 'C': 100, 'M':1000, 'D': 500,
	}

	for i:=0;i<len(s);i++{
		if i < len(s)-1 && romanMap[s[i]] < romanMap[s[i+1]]{
			result -= romanMap[s[i]]
		}else{
			result += romanMap[s[i]]
		}
	} 
	
	return result
}