package gostring

// To export a function in go, name the function to uppercase
func Reverse(word string) string {
	result := ""
	for _, w := range word {
		result = string(w) + result
	}
	return result
}
