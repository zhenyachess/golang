/*
Sample Input:
1 149,6088607594936;1 179,0666666666666

Sample Output:
0.9750
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := strings.Split(strings.ReplaceAll(strings.ReplaceAll(scanner.Text(), " ", ""), ",", "."), ";")
	var num1 float64
	var num2 float64
	num1, _ = strconv.ParseFloat(s[0], 64)
	num2, _ = strconv.ParseFloat(s[1], 64)
	fmt.Printf("%.4f", num1/num2)
}
