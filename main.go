package main

import (
	"fmt"
	"regexp"
	"strings"
)

func replaceNumberToComma(input string) string {
	// Define a regular expression to match "number \n" pattern
	re := regexp.MustCompile(`\n\d+\n`)
	// Replace "number \n" with ","
	replaced := re.ReplaceAllString(input, ", ")

	// Trim any leading or trailing whitespace and remove any empty lines
	replaced = strings.TrimSpace(replaced)
	replaced = strings.ReplaceAll(replaced, "\n\n", "\n")

	return replaced
}

func main() {
	fmt.Println(replaceNumberToComma(`0 - 6 months
Amazon
10
Microsoft
4
BNY Mellon
3
Apple
2
Bloomberg
2
Oracle
2
6 months - 1 year
Goldman Sachs
3
PhonePe
3
EPAM Systems
3
Google
2
1 year - 2 years
Facebook
3
Adobe
3
Uber
2
Walmart Labs
2`))

}
