package main

func romanToInt(s string) int {
	var k1 string
	var k2 string
	result := 0
	m := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IX": 9,
		"IV": 4,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}
	for i := len(s) - 1; i >= 0; i-- {
		k2 = ""
		k1 = s[i : i+1]
		if i >= 1 {
			k2 = s[i-1 : i+1]
		}
		if val, ok := m[k2]; ok {
			i = i - 1
			result = result + val
		} else if val, ok := m[k1]; ok {
			result = result + val
		}
	}
	return result
}

func main() {
	romanToInt("IV")
}
