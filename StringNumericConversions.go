package main

import (
	"fmt"
	s "strings"
)

func menu() {
	fmt.Println("~~~~~~~MENU~~~~~~~")
	fmt.Println("+----------------+")
	fmt.Println("|1. Length       |")
	fmt.Println("|2. Concatenate  |")
	fmt.Println("|3. Uppercase    |")
	fmt.Println("|4. Lowercase    |")
	fmt.Println("|5. Search       |")
	fmt.Println("|6. Split        |")
	fmt.Println("|7. Replace      |")
	fmt.Println("|8. Index        |")
	fmt.Println("|9. Count        |")
	fmt.Println("+----------------+")
}

func main() {
	var option int
	var text, text2, pat string

	fmt.Println("Enter the text : ")
	fmt.Scan(&text)
	menu()
	fmt.Println("Enter the option : ")
	fmt.Scan(&option)
	switch option {
	case 1:
		fmt.Println(len(text))
	case 2:
		fmt.Println("Enter another string to concatenate : ")
		fmt.Scan(&text2)
		s.Join([]string{text, text2}, "*")
	case 3:
		s.ToUpper(text)
	case 4:
		s.ToLower(text)
	case 5:
		fmt.Println("Enter the string to check : ")
		fmt.Scan(&pat)
		s.Contains(text, pat)
	case 6:
		s.Split(text, " ")
	case 7:
		s.Replace("G0", "o", "0", 0)
	case 8:
		fmt.Println("Enter the word to check : ")
		fmt.Scan(&pat)
		s.Index(text, pat)
	case 9:
		fmt.Println("Enter the word to check : ")
		fmt.Scan(&pat)
		s.Count(text, pat)
	}
}
