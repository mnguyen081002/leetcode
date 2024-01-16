package main

func letterCombinations(digits string) []string {
	result := []string{}
	arr := [][]string{{}, {}, {"a", "b", "c"}, {"d", "e", "f"}, {"g", "h", "i"}, {"j", "k", "l"}, {"m", "n", "o"}, {"p", "q", "r", "s"}, {"t", "u", "v"}, {"w", "x", "y", "z"}}

	recursive(digits, "", 0, arr, &result)
	return result
}

func recursive(digits string, cur string, cur_index int, arr [][]string, result *[]string) {
	if cur_index == len(digits) {
		if len(cur) > 0 {
			*result = append(*result, cur)
		}
		return
	}
	di := int(digits[cur_index] - '0')
	for i := 0; i < len(arr[di]); i++ {
		recursive(digits, cur+arr[di][i], cur_index+1, arr, result)
	}
}

func main() {

}
