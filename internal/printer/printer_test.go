package printer_test

import (
	"github.com/olliephillips/pipeline/internal/printer"
	"testing"
)

func TestPrintSalutation(t *testing.T) {
	input := "Tester"
	want := "Hello Tester"
	got := printer.PrintSalutation(input)
	if got != want {
		t.Fatalf("Wanted %s got %s", want, got)
	}
}
