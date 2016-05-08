package nginxconf

import (
	"fmt"
	"testing"
)

var testComments = []string{
	"Foo Bar",
	"### Bar #",
	"こんにちは",
	" #",
}

var expectedText = []string{
	"#Foo Bar",
	"#### Bar #",
	"#こんにちは",
	"# #",
}

func TestComment(t *testing.T) {
	for i, c := range testComments {
		cmt := Comment(c)
		if cmt.String() != expectedText[i] {
			t.Fail()
		}
	}
}

func ExampleComment() {
	c := Comment(" nginxconf comment")
	fmt.Println(c)
	// Output: # nginxconf comment
}
