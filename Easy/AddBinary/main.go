package main

import (
	"fmt"
)

// Ý tưởng:
// 1. Trong nhị phân số lẻ có bit cuối là 1
// 2.

func addBinary(a, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	ib := len(b) - 1

	result := make([]byte, len(a))
	var shifter, sum byte
	for i := len(a) - 1; i >= 0; i-- {
		sum = shifter + a[i] // cộng với bit nhớ
		if ib >= 0 {
			sum += b[ib]
			ib--
		}

		result[i] = sum%2 + '0' // số lẻ thì sẽ là một và cộng với rune '0' trong ascii để về đúng bit. Ví dụ: sum = 98 ('1' + '1') => 0 + 0110000 = 0110000 = '0'
		shifter = sum >> 1 % 2  // dịch bit sang phải 1 để tìm ra có bit nhớ không (nếu 2 số lẻ cộng lại thì bit gần cuối là 1) ví dụ: 98 = 49 + 49 (2 số lẻ)
	}
	if shifter == 0 {
		return string(result)
	}
	return "1" + string(result)
}

func main() {
	fmt.Println(addBinary("1", "11"))
}
