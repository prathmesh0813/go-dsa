package main

func intToRomans(num int) string {
	type romanNumeral struct {
        value int
        symbol string
    }
    var romanNumerals = []romanNumeral{
        {1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
    }

    roman := ""
    for num > 0{
        for _,romanNumeral := range romanNumerals{
            if num >= romanNumeral.value{
                roman += romanNumeral.symbol
                num -= romanNumeral.value 
                break
            }
        }
    }

    return roman
}