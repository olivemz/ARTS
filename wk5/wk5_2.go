func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	commonStr := strs[0]
	var lenStr int
	for _, val := range strs {
		if len(commonStr) == 0 || len(val) == 0 {
			return ""
		} else if len(commonStr) > len(val) {
			lenStr = len(val)
		} else {
			lenStr = len(commonStr)
		}
		for i := 0; i < lenStr; i++ {
			if val[i] != commonStr[i] && i == 0 {
				return ""
			} else if val[i] != commonStr[i] && i > 0 {
				commonStr = val[0:i]
				break
			} else if i == lenStr-1 {
				commonStr = val[0 : i+1]
			}
		}
	}

	return commonStr

}