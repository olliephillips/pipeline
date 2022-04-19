package printer

import "fmt"

// PrintSaluation is simple function that we can unit test, to test the coverage is working
func PrintSalutation(input string) string {
	return fmt.Sprintf("Hello %s", input)
}
