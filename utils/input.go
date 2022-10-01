package utils

import "fmt"

func Input(txt string) string {
	var text string
	fmt.Printf("%s", txt)
	fmt.Scan(&text)
	return text
}
